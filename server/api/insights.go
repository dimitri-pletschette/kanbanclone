package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mattermost/focalboard/server/model"
	"github.com/mattermost/focalboard/server/services/audit"

	mmModel "github.com/mattermost/mattermost-server/v6/model"
)

func (a *API) registerInsightsRoutes(r *mux.Router) {
	// Insights APIs
	r.HandleFunc("/teams/{teamID}/boards/insights", a.sessionRequired(a.handleTeamBoardsInsights)).Methods("GET")
	r.HandleFunc("/users/me/boards/insights", a.sessionRequired(a.handleUserBoardsInsights)).Methods("GET")
}

func (a *API) handleTeamBoardsInsights(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /teams/{teamID}/boards/insights handleTeamBoardsInsights
	//
	// Returns team boards insights
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: teamID
	//   in: path
	//   description: Team ID
	//   required: true
	//   type: string
	// - name: time_range
	//   in: query
	//   description: duration of data to calculate insights for
	//   required: true
	//   type: string
	// - name: page
	//   in: query
	//   description: page offset for top boards
	//   required: true
	//   type: string
	// - name: per_page
	//   in: query
	//   description: limit for boards in a page.
	//   required: true
	//   type: string
	// security:
	// - BearerAuth: []
	// responses:
	//   '200':
	//     description: success
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/BoardInsight"
	//   default:
	//     description: internal error
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"

	if !a.MattermostAuth {
		a.customErrorResponse(w, r.URL.Path, http.StatusNotImplemented, "not permitted in standalone mode", nil)
		return
	}

	vars := mux.Vars(r)
	teamID := vars["teamID"]
	userID := getUserID(r)
	query := r.URL.Query()
	timeRange := query.Get("time_range")

	if pErr := a.ensurePermissionToTeam(userID, teamID, model.PermissionViewTeam); pErr != nil {
		a.errorResponse(w, r, pErr)
		return
	}

	auditRec := a.makeAuditRecord(r, "getTeamBoardsInsights", audit.Fail)
	defer a.audit.LogRecord(audit.LevelRead, auditRec)

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		message := fmt.Sprintf("error converting page parameter to integer: %s", err)
		a.errorResponse(w, r, model.NewErrBadRequest(message))
		return
	}
	if page < 0 {
		a.customErrorResponse(w, r.URL.Path, http.StatusBadRequest, "Invalid page parameter", nil)
	}

	perPage, err := strconv.Atoi(query.Get("per_page"))
	if err != nil {
		message := fmt.Sprintf("error converting per_page parameter to integer: %s", err)
		a.errorResponse(w, r, model.NewErrBadRequest(message))
		return
	}
	if perPage < 0 {
		a.errorResponse(w, r, model.NewErrBadRequest("Invalid page parameter"))
	}

	userTimezone, aErr := a.app.GetUserTimezone(userID)
	if aErr != nil {
		message := fmt.Sprintf("Error getting time zone of user: %s", aErr)
		a.errorResponse(w, r, model.NewErrBadRequest(message))
		return
	}
	userLocation, _ := time.LoadLocation(userTimezone)
	if userLocation == nil {
		userLocation = time.Now().UTC().Location()
	}
	// get unix time for duration
	startTime := mmModel.StartOfDayForTimeRange(timeRange, userLocation)
	boardsInsights, err := a.app.GetTeamBoardsInsights(userID, teamID, &mmModel.InsightsOpts{
		StartUnixMilli: mmModel.GetMillisForTime(*startTime),
		Page:           page,
		PerPage:        perPage,
	})
	if err != nil {
		a.errorResponse(w, r, err)
		return
	}

	data, err := json.Marshal(boardsInsights)
	if err != nil {
		a.errorResponse(w, r, err)
		return
	}

	jsonBytesResponse(w, http.StatusOK, data)

	auditRec.AddMeta("teamBoardsInsightCount", len(boardsInsights.Items))
	auditRec.Success()
}

func (a *API) handleUserBoardsInsights(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /users/me/boards/insights getUserBoardsInsights
	//
	// Returns user boards insights
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: teamID
	//   in: path
	//   description: Team ID
	//   required: true
	//   type: string
	// - name: time_range
	//   in: query
	//   description: duration of data to calculate insights for
	//   required: true
	//   type: string
	// - name: page
	//   in: query
	//   description: page offset for top boards
	//   required: true
	//   type: string
	// - name: per_page
	//   in: query
	//   description: limit for boards in a page.
	//   required: true
	//   type: string
	// security:
	// - BearerAuth: []
	// responses:
	//   '200':
	//     description: success
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/BoardInsight"
	//   default:
	//     description: internal error
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"

	if !a.MattermostAuth {
		a.customErrorResponse(w, r.URL.Path, http.StatusNotImplemented, "not permitted in standalone mode", nil)
		return
	}

	userID := getUserID(r)
	query := r.URL.Query()
	teamID := query.Get("team_id")
	timeRange := query.Get("time_range")

	if pErr := a.ensurePermissionToTeam(userID, teamID, model.PermissionViewTeam); pErr != nil {
		a.errorResponse(w, r, pErr)
		return
	}

	auditRec := a.makeAuditRecord(r, "getUserBoardsInsights", audit.Fail)
	defer a.audit.LogRecord(audit.LevelRead, auditRec)
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		a.errorResponse(w, r, model.NewErrBadRequest("error converting page parameter to integer"))
		return
	}

	if page < 0 {
		a.errorResponse(w, r, model.NewErrBadRequest("Invalid page parameter"))
	}
	perPage, err := strconv.Atoi(query.Get("per_page"))
	if err != nil {
		message := fmt.Sprintf("error converting per_page parameter to integer: %s", err)
		a.errorResponse(w, r, model.NewErrBadRequest(message))
		return
	}

	if perPage < 0 {
		a.errorResponse(w, r, model.NewErrBadRequest("Invalid page parameter"))
	}
	userTimezone, aErr := a.app.GetUserTimezone(userID)
	if aErr != nil {
		message := fmt.Sprintf("Error getting time zone of user: %s", aErr)
		a.errorResponse(w, r, model.NewErrBadRequest(message))
		return
	}
	userLocation, _ := time.LoadLocation(userTimezone)
	if userLocation == nil {
		userLocation = time.Now().UTC().Location()
	}
	// get unix time for duration
	startTime := mmModel.StartOfDayForTimeRange(timeRange, userLocation)
	boardsInsights, err := a.app.GetUserBoardsInsights(userID, teamID, &mmModel.InsightsOpts{
		StartUnixMilli: mmModel.GetMillisForTime(*startTime),
		Page:           page,
		PerPage:        perPage,
	})
	if err != nil {
		a.errorResponse(w, r, err)
		return
	}
	data, err := json.Marshal(boardsInsights)
	if err != nil {
		a.errorResponse(w, r, err)
		return
	}
	jsonBytesResponse(w, http.StatusOK, data)

	auditRec.AddMeta("userBoardInsightCount", len(boardsInsights.Items))
	auditRec.Success()
}
