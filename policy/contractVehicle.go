package policy

import (
	"context"

	"github.com/Capture-411/services-opportunity/services/role"
)

func CanAccessContractVehicle(ctx context.Context) bool {
	return CanCreateContractVehicle(ctx)
}

func CanQueryContractVehicle(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer || r == role.User {
			return true
		}
	}
	return false
}

func CanCreateContractVehicle(ctx context.Context) bool {
	// roles := ExtractRolesClaim(ctx)
	// for _, r := range roles {
	// 	if r == role.Admin || r == role.Reviewer {
	// 		return true
	// 	}
	// }
	return CanQueryContractVehicle(ctx)
}

func CanUpdateContractVehicle(ctx context.Context) bool {
	return CanCreateContractVehicle(ctx)
}

func CanDeleteContractVehicle(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}
