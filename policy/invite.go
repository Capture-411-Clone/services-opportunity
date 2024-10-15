package policy

import (
	"context"

	"github.com/Capture-411/services-opportunity/services/role"
)

func CanAccessInvite(ctx context.Context) bool {
	return CanGetInvite(ctx)
}

func CanGetInvite(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}

func CanQueryInvite(ctx context.Context) bool {
	return CanGetInvite(ctx)
}

func CanCreateInvite(ctx context.Context) bool {
	return CanGetInvite(ctx)
}

func CanUpdateInvite(ctx context.Context) bool {
	return CanGetInvite(ctx)
}

func CanDeleteInvite(ctx context.Context) bool {
	return CanGetInvite(ctx)
}
