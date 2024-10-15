package policy

import (
	"context"
	"fmt"

	"github.com/Capture-411/services-opportunity/kit/jwd"
)

func ExtractRolesClaim(ctx context.Context) []string {
	claims, _ := jwd.Claims(ctx)
	roles := claims["Roles"].([]interface{})

	stringRoles := make([]string, len(roles))

	for _, r := range roles {
		rr := fmt.Sprintf("%v", r)
		stringRoles = append(stringRoles, rr)
	}
	return stringRoles
}

func ExtractIdClaim(ctx context.Context) string {
	claims, _ := jwd.Claims(ctx)
	Id := fmt.Sprintf("%v", claims["Id"])
	return Id
}

func IsAdmin(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, role := range roles {
		if role == "admin" || role == "Admin" {
			return true
		}
	}
	return false
}

func IsReviewer(ctx context.Context) bool {
	roles := ExtractRolesClaim(ctx)
	for _, role := range roles {
		if role == "reviewer" || role == "Reviewer" {
			return true
		}
	}
	return false
}
