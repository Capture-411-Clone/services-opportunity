package policy

import (
	"context"
	"strconv"

	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/role"
)

func CanSeeStaffReport(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}

func CanAccessOpportunity(ctx context.Context) bool {
	return CanQueryOpportunity(ctx)
}

func CanQueryOpportunity(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer || r == role.User {
			return true
		}
	}
	return false
}

func CanCreateOpportunity(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer || r == role.User {
			return true
		}
	}
	return false
}

func CanUpdateOpportunity(ctx context.Context) bool {
	return CanCreateOpportunity(ctx)
}

func CanDeleteOpportunity(ctx context.Context, opportunity models.Opportunity) bool {
	roles := ExtractRolesClaim(ctx)
	Id := ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(Id)
	if int(opportunity.StaffID.Int64) == theID {
		return true
	}

	for _, r := range roles {
		if r == role.Admin || role.Reviewer == r {
			return true
		}
	}
	return false
}

func CanAddOpportunityDocument(ctx context.Context, opportunity models.Opportunity) bool {
	roles := ExtractRolesClaim(ctx)
	Id := ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(Id)

	if opportunity.UserID == theID {
		return true
	}

	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}

func CanApproveOpportunity(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin || r == role.Reviewer {
			return true
		}
	}
	return false
}

func CanImportOpportunity(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	return false
}
