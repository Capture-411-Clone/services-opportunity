package endpoints

import (
	"context"
	"strconv"
	"strings"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

func (s *service) Count(ctx context.Context, params OpportunityQueryRequestParams) (int64, response.ErrorResponse) {
	Id := policy.ExtractIdClaim(ctx)
	userID, _ := strconv.Atoi(Id)

	var count int64

	tx := s.db.WithContext(ctx).Model(&models.Opportunity{})

	where := makeFilters(params)
	if len(where) > 0 {
		tx = tx.Where(strings.Join(where, " AND "))
	}

	if params.MineOnly == "true" {
		tx = tx.Where("opportunities.user_id = ? OR opportunities.staff_id = ?", userID, userID)
	}

	// if only my pipeline: orders.user_id = the user id requesting
	if params.MyPipeline == "true" {
		subQuery := s.db.Model(&models.Order{}).
			Select("opportunity_id").
			Where("orders.user_id = ?", userID).
			Group("opportunity_id")
		tx = tx.Joins("JOIN (?) AS valid_opportunities ON valid_opportunities.opportunity_id = opportunities.id", subQuery)
	}

	err := tx.Count(&count).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return count, response.GormErrorResponse(err, "error in calculate data count")
	}
	return count, response.ErrorResponse{}
}
