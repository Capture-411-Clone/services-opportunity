package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

func (s *service) useInviteCode(ctx context.Context, code string) (models.Invite, response.ErrorResponse) {
	var invite models.Invite
	err := s.db.WithContext(ctx).First(&invite, "code", code).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return invite, response.GormErrorResponse(err, "invite code not found")
	}

	if !(invite.Limit > 0) && invite.Limit != -1 {
		return invite, response.ErrorBadRequest(nil, "invite code expired")
	}
	if invite.Limit == -1 {
		return invite, response.ErrorResponse{}
	}
	if invite.Limit != -1 {
		invite.Limit = invite.Limit - 1
		err = s.db.WithContext(ctx).Save(&invite).Error
		if err != nil {
			s.logger.With(ctx).Error(err)
			return invite, response.GormErrorResponse(err, "error in add invite code")
		}
		return invite, response.ErrorResponse{}
	}
	return invite, response.ErrorResponse{}
}
