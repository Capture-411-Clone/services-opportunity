package endpoints

import (
	"context"
	"fmt"
	"strings"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

func (s *service) Count(ctx context.Context, params AgencyQueryRequestParams) (int64, response.ErrorResponse) {
	var count int64

	tx := s.db.WithContext(ctx).Model(&models.Agency{})

	if params.Query != "" {
		where := query(params)
		for i, w := range where {
			where[i] = fmt.Sprintf("(%s)", w)
		}
		tx = tx.
			Joins("LEFT JOIN departments ON agencies.department_id = departments.id").
			Where(strings.Join(where, " OR "))
	} else {
		where := makeFilters(params)
		for i, w := range where {
			where[i] = fmt.Sprintf("(%s)", w)
		}
		if len(where) > 0 {
			tx = tx.Where(strings.Join(where, " AND "))
		}
	}

	err := tx.Count(&count).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return count, response.GormErrorResponse(err, "error in calculate data count")
	}
	return count, response.ErrorResponse{}
}