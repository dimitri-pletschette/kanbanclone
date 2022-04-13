package sqlstore

import (
	"bytes"
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/mattermost/focalboard/server/utils"

	"strconv"
	"text/template"

	"github.com/mattermost/morph/models"

	"github.com/mattermost/mattermost-server/v6/shared/mlog"

	"github.com/mattermost/morph"
	drivers "github.com/mattermost/morph/drivers"
	mysql "github.com/mattermost/morph/drivers/mysql"
	postgres "github.com/mattermost/morph/drivers/postgres"
	sqlite "github.com/mattermost/morph/drivers/sqlite"
	embedded "github.com/mattermost/morph/sources/embedded"

	mysqldriver "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq" // postgres driver

	sq "github.com/Masterminds/squirrel"

	"github.com/mattermost/focalboard/server/model"
	"github.com/mattermost/mattermost-plugin-api/cluster"
)

//go:embed migrations
var assets embed.FS

const (
	uniqueIDsMigrationRequiredVersion      = 14
	teamsAndBoardsMigrationRequiredVersion = 17

	teamLessBoardsMigrationKey = "TeamLessBoardsMigrationComplete"

	tempSchemaMigrationTableName = "temp_schema_migration"
)

var errChannelCreatorNotInTeam = errors.New("channel creator not found in user teams")

func appendMultipleStatementsFlag(connectionString string) (string, error) {
	config, err := mysqldriver.ParseDSN(connectionString)
	if err != nil {
		return "", err
	}

	if config.Params == nil {
		config.Params = map[string]string{}
	}

	config.Params["multiStatements"] = "true"
	return config.FormatDSN(), nil
}

// migrations in MySQL need to run with the multiStatements flag
// enabled, so this method creates a new connection ensuring that it's
// enabled.
func (s *SQLStore) getMigrationConnection() (*sql.DB, error) {
	connectionString := s.connectionString
	if s.dbType == model.MysqlDBType {
		var err error
		connectionString, err = appendMultipleStatementsFlag(s.connectionString)
		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open(s.dbType, connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *SQLStore) Migrate() error {
	var driver drivers.Driver
	var err error

	migrationConfig := drivers.Config{
		StatementTimeoutInSecs: 1000000,
		MigrationsTable:        fmt.Sprintf("%sschema_migrations", s.tablePrefix),
	}

	if s.dbType == model.SqliteDBType {
		driver, err = sqlite.WithInstance(s.db, &sqlite.Config{Config: migrationConfig})
		if err != nil {
			return err
		}
	}

	var db *sql.DB
	if s.dbType != model.SqliteDBType {
		db, err = s.getMigrationConnection()
		if err != nil {
			return err
		}

		defer db.Close()
	}

	if s.dbType == model.PostgresDBType {
		driver, err = postgres.WithInstance(db, &postgres.Config{Config: migrationConfig})
		if err != nil {
			return err
		}
	}

	if s.dbType == model.MysqlDBType {
		driver, err = mysql.WithInstance(db, &mysql.Config{Config: migrationConfig})
		if err != nil {
			return err
		}
	}

	assetsList, err := assets.ReadDir("migrations")
	if err != nil {
		return err
	}
	assetNamesForDriver := make([]string, len(assetsList))
	for i, dirEntry := range assetsList {
		assetNamesForDriver[i] = dirEntry.Name()
	}

	params := map[string]interface{}{
		"prefix":   s.tablePrefix,
		"postgres": s.dbType == model.PostgresDBType,
		"sqlite":   s.dbType == model.SqliteDBType,
		"mysql":    s.dbType == model.MysqlDBType,
		"plugin":   s.isPlugin,
	}

	migrationAssets := &embedded.AssetSource{
		Names: assetNamesForDriver,
		AssetFunc: func(name string) ([]byte, error) {
			asset, mErr := assets.ReadFile("migrations/" + name)
			if mErr != nil {
				return nil, mErr
			}

			tmpl, pErr := template.New("sql").Parse(string(asset))
			if pErr != nil {
				return nil, pErr
			}
			buffer := bytes.NewBufferString("")

			err = tmpl.Execute(buffer, params)
			if err != nil {
				return nil, err
			}

			return buffer.Bytes(), nil
		},
	}

	src, err := embedded.WithInstance(migrationAssets)
	if err != nil {
		return err
	}

	opts := []morph.EngineOption{
		morph.WithLock("mm-lock-key"),
	}

	if s.dbType == model.SqliteDBType {
		opts = opts[:0] // sqlite driver does not support locking, it doesn't need to anyway.
	}

	engine, err := morph.New(context.Background(), driver, src, opts...)
	if err != nil {
		return err
	}
	defer engine.Close()

	var mutex *cluster.Mutex
	if s.isPlugin {
		var mutexErr error
		mutex, mutexErr = s.NewMutexFn("Boards_dbMutex")
		if mutexErr != nil {
			return fmt.Errorf("error creating database mutex: %w", mutexErr)
		}
	}

	if s.isPlugin {
		s.logger.Debug("Acquiring cluster lock for Unique IDs migration")
		mutex.Lock()
	}

	if err := s.migrateSchemaVersionTable(src.Migrations()); err != nil {
		return err
	}

	if err := ensureMigrationsAppliedUpToVersion(engine, driver, uniqueIDsMigrationRequiredVersion); err != nil {
		return err
	}

	if err := s.runUniqueIDsMigration(); err != nil {
		if s.isPlugin {
			s.logger.Debug("Releasing cluster lock for Unique IDs migration")
			mutex.Unlock()
		}
		return fmt.Errorf("error running unique IDs migration: %w", err)
	}

	if err := ensureMigrationsAppliedUpToVersion(engine, driver, categoriesUUIDIDMigrationRequiredVersion); err != nil {
		return err
	}

	if err := s.runCategoryUUIDIDMigration(); err != nil {
		if s.isPlugin {
			s.logger.Debug("Releasing cluster lock for Unique IDs migration")
			mutex.Unlock()
		}
		return fmt.Errorf("error running categoryID migration: %w", err)
	}

	if err := s.deleteOldSchemaMigrationTable(); err != nil {
		if s.isPlugin {
			mutex.Unlock()
		}
		return err
	}

	if err := ensureMigrationsAppliedUpToVersion(engine, driver, teamsAndBoardsMigrationRequiredVersion); err != nil {
		if s.isPlugin {
			mutex.Unlock()
		}
		return err
	}

	if err := s.migrateTeamLessBoards(); err != nil {
		if s.isPlugin {
			mutex.Unlock()
		}
		return err
	}

	if s.isPlugin {
		s.logger.Debug("Releasing cluster lock for Unique IDs migration")
		mutex.Unlock()
	}

	return engine.ApplyAll()
}

// migrateSchemaVersionTable converts the schema version table from
// the old format used by go-migrate to the new format used by
// gomorph.
// When running the Focalboard with go-migrate's schema version table
// existing in the database, gomorph is unable to make sense of it as it's
// not in the format required by gomorph.
func (s *SQLStore) migrateSchemaVersionTable(migrations []*models.Migration) error {
	migrationNeeded, err := s.isSchemaMigrationNeeded()
	if err != nil {
		return err
	}

	if !migrationNeeded {
		return nil
	}

	s.logger.Info("Migrating schema migration to new format")

	legacySchemaVersion, err := s.getLegacySchemaVersion()
	if err != nil {
		return err
	}

	if err := s.createTempSchemaTable(); err != nil {
		return err
	}

	if err := s.populateTempSchemaTable(migrations, legacySchemaVersion); err != nil {
		return err
	}

	if err := s.useNewSchemaTable(); err != nil {
		return err
	}

	return nil
}

func (s *SQLStore) isSchemaMigrationNeeded() (bool, error) {
	// Check if `dirty` column exists on schema version table.
	// This column exists only for the old schema version table.

	// SQLite needs a bit of a special handling
	if s.dbType == model.SqliteDBType {
		return s.isSchemaMigrationNeededSQLite()
	}

	query := s.getQueryBuilder(s.db).
		Select("count(*)").
		From("information_schema.COLUMNS").
		Where(sq.Eq{
			"TABLE_NAME":  s.tablePrefix + "schema_migrations",
			"COLUMN_NAME": "dirty",
		})

	row := query.QueryRow()

	var count int
	if err := row.Scan(&count); err != nil {
		s.logger.Error("failed to check for columns of schema_migrations table", mlog.Err(err))
		return false, err
	}

	return count == 1, nil
}

func (s *SQLStore) isSchemaMigrationNeededSQLite() (bool, error) {
	// the way to check presence of a column is different
	// for SQLite. Hence, the separate function

	query := fmt.Sprintf("PRAGMA table_info(\"%sschema_migrations\");", s.tablePrefix)
	rows, err := s.db.Query(query)
	if err != nil {
		s.logger.Error("SQLite - failed to check for columns in schema_migrations table", mlog.Err(err))
		return false, err
	}

	defer s.CloseRows(rows)

	data := [][]*string{}
	for rows.Next() {
		// PRAGMA returns 6 columns
		row := make([]*string, 6)

		err := rows.Scan(
			&row[0],
			&row[1],
			&row[2],
			&row[3],
			&row[4],
			&row[5],
		)
		if err != nil {
			s.logger.Error("error scanning rows from SQLite schema_migrations table definition", mlog.Err(err))
			return false, err
		}

		data = append(data, row)
	}

	nameColumnFound := false
	for _, row := range data {
		if len(row) >= 2 && *row[1] == "dirty" {
			nameColumnFound = true
			break
		}
	}

	return nameColumnFound, nil
}

func (s *SQLStore) getLegacySchemaVersion() (uint32, error) {
	query := s.getQueryBuilder(s.db).
		Select("version").
		From(s.tablePrefix + "schema_migrations")

	row := query.QueryRow()

	var version uint32
	if err := row.Scan(&version); err != nil {
		s.logger.Error("error fetching legacy schema version", mlog.Err(err))
		s.logger.Error("getLegacySchemaVersion err " + err.Error())
		return version, err
	}

	return version, nil
}

func (s *SQLStore) createTempSchemaTable() error {
	// squirrel doesn't support DDL query in query builder
	// so, we need to use a plain old string
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (Version bigint NOT NULL, Name varchar(64) NOT NULL, PRIMARY KEY (Version))", s.tablePrefix+tempSchemaMigrationTableName)
	if _, err := s.db.Exec(query); err != nil {
		s.logger.Error("failed to create temporary schema migration table", mlog.Err(err))
		s.logger.Error("createTempSchemaTable error  " + err.Error())
		return err
	}

	return nil
}
func (s *SQLStore) populateTempSchemaTable(migrations []*models.Migration, legacySchemaVersion uint32) error {
	query := s.getQueryBuilder(s.db).
		Insert(s.tablePrefix+tempSchemaMigrationTableName).
		Columns("Version", "Name")

	for _, migration := range migrations {
		// migrations param contains both up and down variant for
		// each migration. Skipping for either one (down in this case)
		// to process a migration only a single time.
		if migration.Direction == models.Down {
			continue
		}

		if migration.Version > legacySchemaVersion {
			break
		}

		query = query.Values(migration.Version, migration.Name)
	}

	if _, err := query.Exec(); err != nil {
		s.logger.Error("failed to insert migration records into temporary schema table", mlog.Err(err))
		return err
	}

	return nil
}

func (s *SQLStore) useNewSchemaTable() error {
	// first delete the old table, then
	// rename the new table to old table's name

	// renaming old schema migration table. Will delete later once the migration is
	// complete, just in case.
	var query string
	if s.dbType == model.MysqlDBType {
		query = fmt.Sprintf("RENAME TABLE `%sschema_migrations` TO `%sschema_migrations_old_temp`", s.tablePrefix, s.tablePrefix)
	} else {
		query = fmt.Sprintf("ALTER TABLE %sschema_migrations RENAME TO %sschema_migrations_old_temp", s.tablePrefix, s.tablePrefix)
	}

	if _, err := s.db.Exec(query); err != nil {
		s.logger.Error("failed to rename old schema migration table", mlog.Err(err))
		return err
	}

	// renaming new temp table to old table's name
	if s.dbType == model.MysqlDBType {
		query = fmt.Sprintf("RENAME TABLE `%s%s` TO `%sschema_migrations`", s.tablePrefix, tempSchemaMigrationTableName, s.tablePrefix)
	} else {
		query = fmt.Sprintf("ALTER TABLE %s%s RENAME TO %sschema_migrations", s.tablePrefix, tempSchemaMigrationTableName, s.tablePrefix)
	}

	if _, err := s.db.Exec(query); err != nil {
		s.logger.Error("failed to rename temp schema table", mlog.Err(err))
		return err
	}

	return nil
}

func (s *SQLStore) deleteOldSchemaMigrationTable() error {
	query := "DROP TABLE IF EXISTS " + s.tablePrefix + "schema_migrations_old_temp"
	if _, err := s.db.Exec(query); err != nil {
		s.logger.Error("failed to delete old temp schema migrations table", mlog.Err(err))
		return err
	}

	return nil
}

// We no longer support boards existing in DMs and private
// group messages. This function migrates all boards
// belonging to a DM to the best possible team.
func (s *SQLStore) migrateTeamLessBoards() error {
	if !s.isPlugin {
		return nil
	}

	setting, err := s.GetSystemSetting(teamLessBoardsMigrationKey)
	if err != nil {
		return fmt.Errorf("cannot get teamless boards migration state: %w", err)
	}

	// If the migration is already completed, do not run it again.
	if hasAlreadyRun, _ := strconv.ParseBool(setting); hasAlreadyRun {
		return nil
	}

	boards, err := s.getDMBoards(s.db)
	if err != nil {
		return err
	}

	s.logger.Info(fmt.Sprintf("Migrating %d teamless boards to a team", len(boards)))

	// cache for best suitable team for a DM. Since a DM can
	// contain multiple boards, caching this avoids
	// duplicate queries for the same DM.
	channelToTeamCache := map[string]string{}

	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		s.logger.Error("error starting transaction in migrateTeamLessBoards", mlog.Err(err))
		return err
	}

	for i := range boards {
		// check the cache first
		teamID, ok := channelToTeamCache[boards[i].ChannelID]

		// query DB if entry not found in cache
		if !ok {
			teamID, err = s.getBestTeamForBoard(s.db, boards[i])
			if err != nil {
				// don't let one board's error spoil
				// the mood for others
				continue
			}
		}

		channelToTeamCache[boards[i].ChannelID] = teamID
		boards[i].TeamID = teamID

		query := s.getQueryBuilder(tx).
			Update(s.tablePrefix+"boards").
			Set("team_id", teamID).
			Set("type", model.BoardTypePrivate).
			Where(sq.Eq{"id": boards[i].ID})

		if _, err := query.Exec(); err != nil {
			s.logger.Error("failed to set team id for board", mlog.String("board_id", boards[i].ID), mlog.String("team_id", teamID), mlog.Err(err))
			return err
		}
	}

	if err := s.setSystemSetting(tx, teamLessBoardsMigrationKey, strconv.FormatBool(true)); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "migrateTeamLessBoards"))
		}
		return fmt.Errorf("cannot mark migration as completed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		s.logger.Error("failed to commit migrateTeamLessBoards transaction", mlog.Err(err))
		return err
	}

	return nil
}

func (s *SQLStore) getDMBoards(tx sq.BaseRunner) ([]*model.Board, error) {
	conditions := sq.And{
		sq.Eq{"team_id": ""},
		sq.Or{
			sq.Eq{"type": "D"},
			sq.Eq{"type": "G"},
		},
	}

	boards, err := s.getBoardsByCondition(tx, conditions)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return []*model.Board{}, nil
	}

	return boards, err
}

// The destination is selected as the first team where all members
// of the DM are a part of. If no such team exists,
// we use the first team to which DM creator belongs to.
func (s *SQLStore) getBestTeamForBoard(tx sq.BaseRunner, board *model.Board) (string, error) {
	userTeams, err := s.getBoardUserTeams(tx, board)
	if err != nil {
		return "", err
	}

	teams := [][]interface{}{}
	for _, userTeam := range userTeams {
		userTeamInterfaces := make([]interface{}, len(userTeam))
		for i := range userTeam {
			userTeamInterfaces[i] = userTeam[i]
		}
		teams = append(teams, userTeamInterfaces)
	}

	commonTeams := utils.Intersection(teams...)
	var teamID string
	if len(commonTeams) > 0 {
		teamID = commonTeams[0].(string)
	} else {
		// no common teams found. Let's try finding the best suitable team
		if board.Type == "D" {
			// get DM's creator and pick one of their team
			channel, appErr := (*s.pluginAPI).GetChannel(board.ChannelID)
			if appErr != nil {
				s.logger.Error("failed to fetch DM channel for board", mlog.String("board_id", board.ID), mlog.String("channel_id", board.ChannelID), mlog.Err(appErr))
				return "", appErr
			}

			if _, ok := userTeams[channel.CreatorId]; !ok {
				err := fmt.Errorf("%w board_id: %s, channel_id: %s, creator_id: %s", errChannelCreatorNotInTeam, board.ID, board.ChannelID, channel.CreatorId)
				s.logger.Error(err.Error())
				return "", err
			}

			teamID = userTeams[channel.CreatorId][0]
		} else if board.Type == "G" {
			// pick the team that has the most users as members
			teamFrequency := map[string]int{}
			highestFrequencyTeam := ""
			highestFrequencyTeamFrequency := -1

			for _, teams := range userTeams {
				for _, teamID := range teams {
					teamFrequency[teamID]++

					if teamFrequency[teamID] > highestFrequencyTeamFrequency {
						highestFrequencyTeamFrequency = teamFrequency[teamID]
						highestFrequencyTeam = teamID
					}
				}
			}

			teamID = highestFrequencyTeam
		}
	}

	return teamID, nil
}

func (s *SQLStore) getBoardUserTeams(tx sq.BaseRunner, board *model.Board) (map[string][]string, error) {
	query := s.getQueryBuilder(tx).
		Select("teammembers.userid", "teammembers.teamid").
		From("channelmembers").
		Join("teammembers ON channelmembers.userid = teammembers.userid").
		Where(sq.Eq{"channelid": board.ChannelID})

	rows, err := query.Query()
	if err != nil {
		s.logger.Error("failed to fetch user teams for board", mlog.String("boardID", board.ID), mlog.String("channelID", board.ChannelID), mlog.Err(err))
		return nil, err
	}

	defer rows.Close()

	userTeams := map[string][]string{}

	for rows.Next() {
		var userID, teamID string
		err := rows.Scan(&userID, &teamID)
		if err != nil {
			s.logger.Error("getBoardUserTeams failed to scan SQL query result", mlog.String("boardID", board.ID), mlog.String("channelID", board.ChannelID), mlog.Err(err))
			return nil, err
		}

		userTeams[userID] = append(userTeams[userID], teamID)
	}

	return userTeams, nil
}

func ensureMigrationsAppliedUpToVersion(engine *morph.Morph, driver drivers.Driver, version int) error {
	applied, err := driver.AppliedMigrations()
	if err != nil {
		return err
	}
	currentVersion := len(applied)

	// if the target version is below or equal to the current one, do
	// not migrate either because is not needed (both are equal) or
	// because it would downgrade the database (is below)
	if version <= currentVersion {
		return nil
	}

	if _, err = engine.Apply(version - currentVersion); err != nil {
		return err
	}

	return nil
}
