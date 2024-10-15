package policy

import (
	"context"
	"strconv"

	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/role"
)

func CanQueryAllNotifications(ctx context.Context) bool {
	return CanCreateNotification(ctx)
}

func CanCreateNotification(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}

func CanUpdateNotification(ctx context.Context, notification models.Notification) bool {
	roles := ExtractRolesClaim(ctx)
	Id := ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(Id)

	if *notification.TargetUserID == uint(theID) {
		return true
	}

	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}

func CanDeleteNotification(ctx context.Context, notification models.Notification) bool {
	return CanUpdateNotification(ctx, notification)
}
