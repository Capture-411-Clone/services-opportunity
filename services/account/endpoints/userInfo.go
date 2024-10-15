package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

func (s *service) UserInfo(ctx context.Context, accessToken string) (models.User, response.ErrorResponse) {
	Id := policy.ExtractIdClaim(ctx)
	var user models.User

	err := s.db.WithContext(ctx).
		Preload("Roles").
		First(&user, "id", Id).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.User{}, response.ErrorUnAuthorized(nil, "cannot find a user with these credentials")
	}

	if user.SuspendedAt.Valid {
		s.logger.With(ctx).Error(err)
		return user, response.ErrorUnAuthorized(nil, "your account has been suspended")
	}

	return user, response.ErrorResponse{}
}
