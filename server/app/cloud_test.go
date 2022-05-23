package app

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/mattermost/focalboard/server/model"

	mmModel "github.com/mattermost/mattermost-server/v6/model"
)

func TestIsCloud(t *testing.T) {
	t.Run("if it's not running on plugin mode", func(t *testing.T) {
		th, tearDown := SetupTestHelper(t)
		defer tearDown()

		th.Store.EXPECT().GetLicense().Return(nil)
		require.False(t, th.App.IsCloud())
	})

	t.Run("if it's running on plugin mode but the license is incomplete", func(t *testing.T) {
		th, tearDown := SetupTestHelper(t)
		defer tearDown()

		fakeLicense := &mmModel.License{}

		th.Store.EXPECT().GetLicense().Return(fakeLicense)
		require.False(t, th.App.IsCloud())

		fakeLicense = &mmModel.License{Features: &mmModel.Features{}}

		th.Store.EXPECT().GetLicense().Return(fakeLicense)
		require.False(t, th.App.IsCloud())
	})

	t.Run("if it's running on plugin mode, with a non-cloud license", func(t *testing.T) {
		th, tearDown := SetupTestHelper(t)
		defer tearDown()

		fakeLicense := &mmModel.License{
			Features: &mmModel.Features{Cloud: mmModel.NewBool(false)},
		}

		th.Store.EXPECT().GetLicense().Return(fakeLicense)
		require.False(t, th.App.IsCloud())
	})

	t.Run("if it's running on plugin mode with a cloud license", func(t *testing.T) {
		th, tearDown := SetupTestHelper(t)
		defer tearDown()

		fakeLicense := &mmModel.License{
			Features: &mmModel.Features{Cloud: mmModel.NewBool(true)},
		}

		th.Store.EXPECT().GetLicense().Return(fakeLicense)
		require.True(t, th.App.IsCloud())
	})
}

func TestIsCloudLimited(t *testing.T) {
	t.Run("if no limit has been set, it should be false", func(t *testing.T) {
		th, tearDown := SetupTestHelper(t)
		defer tearDown()

		require.Zero(t, th.App.CardLimit)
		require.False(t, th.App.IsCloudLimited())
	})

	t.Run("if the limit is set, it should be true", func(t *testing.T) {
		th, tearDown := SetupTestHelper(t)
		defer tearDown()

		fakeLicense := &mmModel.License{
			Features: &mmModel.Features{Cloud: mmModel.NewBool(true)},
		}
		th.Store.EXPECT().GetLicense().Return(fakeLicense)

		th.App.CardLimit = 5
		require.True(t, th.App.IsCloudLimited())
	})
}

func TestSetCloudLimits(t *testing.T) {
	t.Run("if the limits are empty, it should do nothing", func(t *testing.T) {
		t.Run("limits empty", func(t *testing.T) {
			th, tearDown := SetupTestHelper(t)
			defer tearDown()

			require.Zero(t, th.App.CardLimit)

			require.NoError(t, th.App.SetCloudLimits(nil))
			require.Zero(t, th.App.CardLimit)
		})

		t.Run("limits not empty but board limits empty", func(t *testing.T) {
			th, tearDown := SetupTestHelper(t)
			defer tearDown()

			require.Zero(t, th.App.CardLimit)

			limits := &mmModel.ProductLimits{}

			require.NoError(t, th.App.SetCloudLimits(limits))
			require.Zero(t, th.App.CardLimit)
		})
	})

	t.Run("if the limits are not empty, it should update them and calculate the new timestamp", func(t *testing.T) {
		th, tearDown := SetupTestHelper(t)
		defer tearDown()

		require.Zero(t, th.App.CardLimit)

		newCardLimitTimestamp := int64(27)
		th.Store.EXPECT().UpdateCardLimitTimestamp(5).Return(newCardLimitTimestamp, nil)

		limits := &mmModel.ProductLimits{
			Boards: &mmModel.BoardsLimits{Cards: mmModel.NewInt(5)},
		}

		require.NoError(t, th.App.SetCloudLimits(limits))
		require.Equal(t, 5, th.App.CardLimit)
	})

	t.Run("if the limits are already set and we unset them, the timestamp will be unset too", func(t *testing.T) {
		th, tearDown := SetupTestHelper(t)
		defer tearDown()

		th.App.CardLimit = 20

		th.Store.EXPECT().UpdateCardLimitTimestamp(0)

		require.NoError(t, th.App.SetCloudLimits(nil))

		require.Zero(t, th.App.CardLimit)
	})

	t.Run("if the limits are already set and we try to set the same ones again", func(t *testing.T) {
		th, tearDown := SetupTestHelper(t)
		defer tearDown()

		th.App.CardLimit = 20

		// the call to update card limit timestamp should not happen
		// as the limits didn't change
		th.Store.EXPECT().UpdateCardLimitTimestamp(gomock.Any()).Times(0)

		limits := &mmModel.ProductLimits{
			Boards: &mmModel.BoardsLimits{Cards: mmModel.NewInt(20)},
		}

		require.NoError(t, th.App.SetCloudLimits(limits))
		require.Equal(t, 20, th.App.CardLimit)
	})
}

func TestUpdateCardLimitTimestamp(t *testing.T) {
	fakeLicense := &mmModel.License{
		Features: &mmModel.Features{Cloud: mmModel.NewBool(true)},
	}

	t.Run("if the server is a cloud instance but not limited, it should do nothing", func(t *testing.T) {
		th, tearDown := SetupTestHelper(t)
		defer tearDown()

		require.Zero(t, th.App.CardLimit)

		// the license check will not be done as the limit not being
		// set is enough for the method to return
		th.Store.EXPECT().GetLicense().Times(0)
		// no call to UpdateCardLimitTimestamp should happen as the
		// method should shortcircuit if not cloud limited
		th.Store.EXPECT().UpdateCardLimitTimestamp(gomock.Any()).Times(0)

		require.NoError(t, th.App.UpdateCardLimitTimestamp())
	})

	t.Run("if the server is a cloud instance and the timestamp is set, it should run the update", func(t *testing.T) {
		th, tearDown := SetupTestHelper(t)
		defer tearDown()

		th.App.CardLimit = 5

		th.Store.EXPECT().GetLicense().Return(fakeLicense)
		// no call to UpdateCardLimitTimestamp should happen as the
		// method should shortcircuit if not cloud limited
		th.Store.EXPECT().UpdateCardLimitTimestamp(5)

		require.NoError(t, th.App.UpdateCardLimitTimestamp())
	})
}
