package policy

import (
	"context"

	"github.com/Capture-411/services-opportunity/services/role"
)

func CanAccessNaics(ctx context.Context) bool {
	return CanCreateNaics(ctx)
}

func CanQueryNaics(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer || r == role.User {
			return true
		}
	}
	return false
}

func CanCreateNaics(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer {
			return true
		}
	}
	return false
}

func CanUpdateNaics(ctx context.Context) bool {
	return CanCreateNaics(ctx)
}

func CanDeleteNaics(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}
