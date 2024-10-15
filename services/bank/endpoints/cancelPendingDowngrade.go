package endpoints

import (
	"context"
	"strconv"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

type CancelPendingDowngradeRequestParams struct {
	DowngradeID string `json:"downgrade_id" openapi:"example:PbRzzvjoOCsfMQ;nullable;in:path"`
}

func (p *CancelPendingDowngradeRequestParams) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"downgrade_id": gocriteria.New("downgrade_id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"downgrade_id": "Downgrade ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) CancelPendingDowngrade(ctx context.Context, params CancelPendingDowngradeRequestParams) response.ErrorResponse {
	userid := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(userid)

	// Fetch user by userID
	var user models.User
	err := s.db.Model(&models.User{}).Where("id = ?", theID).First(&user).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error fetching user: %v", err)
		return response.ErrorBadRequest(nil, "Could not find user")
	}

	var downgrade models.Downgrade
	err = s.db.Model(&models.Downgrade{}).Where("id = ?", params.DowngradeID).First(&downgrade).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error fetching downgrade: %v", err)
		return response.ErrorBadRequest(nil, "Could not find downgrade")
	}

	// Database Downgrade Model operation
	err = s.db.Model(&models.Downgrade{}).Where("id = ?", params.DowngradeID).Delete(&downgrade).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error deleting downgrade: %v", err)
		return response.ErrorBadRequest(nil, "Could not cancel downgrade")
	}

	go s.sendCancelDowngradeNotification(ctx, user)

	return response.ErrorResponse{} // Return success response
}

func (s *service) sendCancelDowngradeNotification(ctx context.Context, user models.User) {
	var targetUsers []models.User
	targetUsers = append(targetUsers, user)
	// Send notification for user
	s.notifier.SendNotification(ctx, targetUsers, nil, []string{"email"}, "Your subscription has been downgraded", "", "Subscription Downgraded")
}
