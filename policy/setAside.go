package policy

import (
	"context"

	"github.com/Capture-411/services-opportunity/services/role"
)

func CanAccessSetAside(ctx context.Context) bool {
	return CanCreateSetAside(ctx)
}

func CanQuerySetAside(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer || r == role.User {
			return true
		}
	}
	return false
}

func CanCreateSetAside(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer {
			return true
		}
	}
	return false
}

func CanUpdateSetAside(ctx context.Context) bool {
	return CanCreateSetAside(ctx)
}

func CanDeleteSetAside(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}
