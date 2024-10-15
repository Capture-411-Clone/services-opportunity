package endpoints

import (
	"context"
	"strings"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

func (s *service) Count(ctx context.Context, params OrderQueryRequestParams) (int64, response.ErrorResponse) {
	var count int64

	where := makeFilters(params)

	tx := s.db.WithContext(ctx).Model(&models.Order{})

	if len(where) > 0 {
		tx = tx.Where(strings.Join(where, " AND "))
	}

	err := tx.Count(&count).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return count, response.GormErrorResponse(err, "error in calculate data count")
	}
	return count, response.ErrorResponse{}
}
