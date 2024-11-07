package endpoints

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/Capture-411/services-opportunity/kit/jwd"
	"github.com/golang-jwt/jwt/v4"
)

func (s *service) generateAccessTokenJWT(identity Identity) (string, error) {
	// Create header with HMAC-SHA256 algorithm and encryption flag
	header := jwd.JWTHeader{
		Algorithm: "HS256",
		Type:      "JWT",
		Encrypted: true,
	}
	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	// Create claims
	claims := jwt.MapClaims{
		"Id":        identity.GetID(),
		"Roles":     []interface{}{},
		"ExpiresAt": time.Now().Add(time.Duration(s.tokenExpiration) * time.Minute).Unix(),
		"IssuedAt":  time.Now().Unix(),
		"Nonce":     jwd.GenerateNonce(),
	}

	if len(identity.GetRoles()) > 0 {
		claims["Roles"] = identity.GetRoles()
	}

	// Marshal claims
	payloadBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	// Derive encryption key from signing key
	encKey := sha256.Sum256([]byte(s.accessTokenSigningKey))

	// Encrypt payload
	encryptedPayload, err := jwd.Encrypt(payloadBytes, encKey[:])
	if err != nil {
		return "", err
	}

	// Encode header and encrypted payload
	encodedHeader := base64.RawURLEncoding.EncodeToString(headerBytes)
	encodedPayload := base64.RawURLEncoding.EncodeToString(encryptedPayload)

	// Create signing input
	signingInput := encodedHeader + "." + encodedPayload

	// Generate HMAC-SHA256 signature
	signature := jwd.GenerateHMACSignature([]byte(signingInput), []byte(s.accessTokenSigningKey))

	// Create final token
	token := signingInput + "." + base64.RawURLEncoding.EncodeToString(signature)

	return token, nil
}
