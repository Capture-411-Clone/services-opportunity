package endpoints

import (
	"time"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/golang-jwt/jwt"
)

func (s *service) generateAccessTokenJWT(identity Identity) (string, error) {
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"Audience":  config.AppConfig.HostUri,
			"ExpiresAt": time.Now().Add(time.Duration(s.tokenExpiration) * time.Minute).Unix(),
			"Id":        identity.GetID(),
			"IssuedAt":  time.Now().Unix(),
			"Issuer":    identity.GetFullName(),
			"NotBefore": time.Now().Unix(),
			"Subject":   identity.GetPhone(),
			"Roles":     identity.GetRoles(),
		},
	).SignedString([]byte(s.accessTokenSigningKey))
	if err != nil {
		s.logger.Error("sign access token", err)
		return "", err
	}
	return token, err
}
