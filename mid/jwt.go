package mid

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/kit/jwd"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func EchoJWTHandler() echo.MiddlewareFunc {
	//lint:ignore SA1019 ignore deprecated warning for staticcheck
	return middleware.JWTWithConfig(
		middleware.JWTConfig{
			ContextKey:  "user",
			SigningKey:  []byte(config.AppConfig.AccessTokenSigningKey),
			TokenLookup: "header:Authorization",
			AuthScheme:  "Bearer",
			ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
				parts := strings.Split(auth, ".")
				if len(parts) != 3 {
					return nil, errors.New("invalid token format")
				}

				// Decode header
				headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
				if err != nil {
					return nil, errors.New("invalid header encoding")
				}

				var header jwd.JWTHeader
				if err := json.Unmarshal(headerBytes, &header); err != nil {
					return nil, errors.New("invalid header format")
				}

				if header.Algorithm != "HS256" {
					return nil, errors.New("unsupported signing algorithm")
				}

				// Verify signature first
				signature, err := base64.RawURLEncoding.DecodeString(parts[2])
				if err != nil {
					return nil, errors.New("invalid signature encoding")
				}

				signingInput := parts[0] + "." + parts[1]
				if !jwd.VerifyHMACSignature(
					[]byte(signingInput),
					[]byte(config.AppConfig.AccessTokenSigningKey),
					signature,
				) {
					return nil, errors.New("invalid signature")
				}

				// Decode encrypted payload
				encryptedPayload, err := base64.RawURLEncoding.DecodeString(parts[1])
				if err != nil {
					return nil, errors.New("invalid payload encoding")
				}

				// Derive encryption key
				encKey := sha256.Sum256([]byte(config.AppConfig.AccessTokenSigningKey))

				// Decrypt payload
				payloadBytes, err := jwd.Decrypt(encryptedPayload, encKey[:])
				if err != nil {
					return nil, errors.New("failed to decrypt payload")
				}

				var claims jwt.MapClaims
				if err := json.Unmarshal(payloadBytes, &claims); err != nil {
					return nil, errors.New("invalid claims format")
				}

				// Check expiration
				if exp, ok := claims["ExpiresAt"].(float64); ok {
					if int64(exp) < time.Now().Unix() {
						return nil, errors.New("token has expired")
					}
				}

				// Ensure Roles is initialized
				if claims["Roles"] == nil {
					claims["Roles"] = []interface{}{}
				}

				token := &jwt.Token{
					Raw:       auth,
					Claims:    claims,
					Valid:     true,
					Signature: parts[2],
					Method:    jwt.SigningMethodHS256,
					Header:    map[string]interface{}{"alg": "HS256", "typ": "JWT", "enc": true},
				}

				return token, nil
			},
		},
	)
}

func BindUserToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		token, ok := c.Get("user").(*jwt.Token)
		if !ok || token == nil {
			return errors.New("invalid or missing token")
		}

		// Ensure token claims are properly initialized
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["Roles"] == nil {
				claims["Roles"] = []interface{}{}
			}
		}

		ctx = context.WithValue(ctx, jwd.UserContextKey, token)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
