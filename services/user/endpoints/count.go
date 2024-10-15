package endpoints

import (
	"context"
	"strings"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

func (s *service) Count(ctx context.Context, params UserQueryRequestParams) (int64, response.ErrorResponse) {
	var count int64

	tx := s.db.WithContext(ctx).Model(&models.User{})

	where := makeFilters(params)
	if len(where) > 0 {
		tx = tx.Where(strings.Join(where, " AND "))
	}

	err := tx.Model(&models.User{}).Count(&count).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return count, response.GormErrorResponse(err, "خطا در محاسبه داده")
	}

	return count, response.ErrorResponse{}
}
