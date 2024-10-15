package endpoints

import (
	"context"
	"strconv"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
	"github.com/Capture-411/services-opportunity/utils"
)

func (s *service) GetAllOpportunityOwnedIDs(ctx context.Context) ([]int64, response.ErrorResponse) {
	Id := policy.ExtractIdClaim(ctx)
	userID, _ := strconv.Atoi(Id)

	var orders []models.Order

	err := s.db.Model(models.Order{}).Where("user_id = ? AND paid_at IS NOT NULL AND refunded_at IS NULL", userID).Find(&orders).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error in finding orders: %v", err)
		return []int64{}, response.GormErrorResponse(err, "error in finding orders")
	}

	var ownedIDs []int64
	for _, order := range orders {
		// if order.OpportunityID already in owned escape it
		if utils.Contains(ownedIDs, order.OpportunityID) {
			continue
		}
		ownedIDs = append(ownedIDs, order.OpportunityID)
	}

	return ownedIDs, response.ErrorResponse{}
}
