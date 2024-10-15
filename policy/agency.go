package policy

import (
	"context"

	"github.com/Capture-411/services-opportunity/services/role"
)

func CanAccessAgency(ctx context.Context) bool {
	return CanCreateAgency(ctx)
}

func CanQueryAgency(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer || r == role.User {
			return true
		}
	}
	return false
}

func CanCreateAgency(ctx context.Context) bool {
	// roles := ExtractRolesClaim(ctx)
	// for _, r := range roles {
	// 	if r == role.Admin || r == role.Reviewer {
	// 		return true
	// 	}
	// }
	return CanQueryAgency(ctx)
}

func CanUpdateAgency(ctx context.Context) bool {
	return CanCreateAgency(ctx)
}

func CanDeleteAgency(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}
