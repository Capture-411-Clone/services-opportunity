package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

func (s *service) Info(ctx context.Context) (models.SiteInfo, response.ErrorResponse) {
	var site models.SiteInfo
	err := s.db.WithContext(ctx).First(&site, "id", 1).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return site, response.GormErrorResponse(err, "invite code not found")
	}

	return site, response.ErrorResponse{}
}
