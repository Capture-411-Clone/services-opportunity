package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

func (s *service) Use(ctx context.Context, code string) (models.Invite, response.ErrorResponse) {
	var invite models.Invite
	err := s.db.WithContext(ctx).First(&invite, "code", code).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return invite, response.GormErrorResponse(nil, "invitation code not found")
	}

	if !(invite.Limit > 0) || invite.Limit != -1 {
		return invite, response.GormErrorResponse(err, "your invite code has been used")
	}
	if invite.Limit == -1 {
		return invite, response.ErrorResponse{}
	}
	if invite.Limit != -1 {
		invite.Limit = invite.Limit - 1
		err := s.db.WithContext(ctx).Save(&invite).Error
		if err != nil {
			s.logger.With(ctx).Error(err)
			return invite, response.GormErrorResponse(err, "an error occurred while using the invite code")
		}
	}
	return invite, response.ErrorResponse{}
}
