package endpoints

import (
	"context"
	"strings"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

func (s *service) Count(ctx context.Context, requestParams NotificationQueryRequestParams) (int64, response.ErrorResponse) {
	var count int64

	where := makeFilters(requestParams)

	tx := s.db.WithContext(ctx).Model(&models.Notification{})

	if len(where) > 0 {
		tx = tx.Where(strings.Join(where, " AND "))
	}

	err := tx.Count(&count).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return count, response.GormErrorResponse(err, "Error in data query")
	}

	return count, response.ErrorResponse{}
}
