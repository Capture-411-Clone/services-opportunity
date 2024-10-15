package policy

import (
	"context"

	"github.com/Capture-411/services-opportunity/services/role"
)

func CanAccessOrder(ctx context.Context) bool {
	return CanQueryOrder(ctx)
}

func CanToggleRefundAt(ctx context.Context) bool {
	return CanDeleteOrder(ctx)
}

func CanQueryOrder(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer {
			return true
		}
	}
	return false
}

func CanDeleteOrder(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}
