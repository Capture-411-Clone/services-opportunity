package policy

import (
	"context"

	"github.com/Capture-411/services-opportunity/services/role"
)

func CanAccessOffice(ctx context.Context) bool {
	return CanCreateOffice(ctx)
}

func CanQueryOffice(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer || r == role.User {
			return true
		}
	}
	return false
}

func CanCreateOffice(ctx context.Context) bool {
	// roles := ExtractRolesClaim(ctx)
	// for _, r := range roles {
	// 	if r == role.Admin || r == role.Reviewer {
	// 		return true
	// 	}
	// }
	return CanQueryOffice(ctx)
}

func CanUpdateOffice(ctx context.Context) bool {
	return CanCreateOffice(ctx)
}

func CanDeleteOffice(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}
