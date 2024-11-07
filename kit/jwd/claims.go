package jwd

import (
	"context"
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

func Claims(ctx context.Context) (jwt.MapClaims, error) {
	token, ok := ctx.Value(UserContextKey).(*jwt.Token)
	if !ok || token == nil {
		log.Println("No valid JWT token found in context (Claims method)")
		return jwt.MapClaims{"Roles": []interface{}{}}, nil // Return empty roles instead of nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("conversion to jwt.MapClaims failed for your token in context (Claims method)")
	}

	// Ensure Roles exists and is initialized
	if claims["Roles"] == nil {
		claims["Roles"] = []interface{}{}
	}

	return claims, nil
}
