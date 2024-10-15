package policy

import (
	"context"

	"github.com/Capture-411/services-opportunity/services/role"
)

func CanAccessMarket(ctx context.Context) bool {
	return CanCreateMarket(ctx)
}

func CanQueryMarket(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer || r == role.User {
			return true
		}
	}
	return false
}

func CanCreateMarket(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer {
			return true
		}
	}
	return false
}

func CanUpdateMarket(ctx context.Context) bool {
	return CanCreateMarket(ctx)
}

func CanDeleteMarket(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}
