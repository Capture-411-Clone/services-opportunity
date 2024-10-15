package policy

import (
	"context"
	"strconv"

	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/role"
)

func CanUpdateDocument(ctx context.Context, document models.Document) bool {
	roles := ExtractRolesClaim(ctx)
	Id := ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(Id)
	for _, r := range roles {
		if r == role.Admin || document.UserID == theID {
			return true
		}
	}
	return false
}

func CanDeleteDocument(ctx context.Context, documents []models.Document) bool {
	roles := ExtractRolesClaim(ctx)
	Id := ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(Id)
	for _, r := range roles {
		if r == role.Admin {
			return true
		}
	}
	var allBelongToUser bool
	for _, m := range documents {
		allBelongToUser = m.UserID == theID
		if !allBelongToUser {
			break
		}
	}
	return allBelongToUser
}

func CanAccessDocument(ctx context.Context, doc models.Document, user models.User) bool {
	if IsAdmin(ctx) || IsReviewer(ctx) || user.ID == uint(doc.UserID) || user.ID == uint(doc.StaffID.Int64) {
		return true
	}
	return false
}
