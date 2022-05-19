// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package notifysubscriptions

import (
	"github.com/mattermost/focalboard/server/model"

	mm_model "github.com/mattermost/mattermost-server/v6/model"
)

// SubscriptionDelivery provides an interface for delivering subscription notifications to other systems, such as
// channels server via plugin API.
type SubscriptionDelivery interface {
	SubscriptionDeliverSlackAttachments(workspaceID string, teamID string, subscriberID string, subscriberType model.SubscriberType,
		attachments []*mm_model.SlackAttachment) error

	GetTeamIDForWorkspace(workspaceID string) (string, error)
}
