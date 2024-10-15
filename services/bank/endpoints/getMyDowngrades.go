package endpoints

import (
	"context"
	"strconv"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

func (s *service) GetMyDowngrades(ctx context.Context) ([]models.Downgrade, response.ErrorResponse) {
	userid := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(userid)

	// Fetch user by userID
	var user models.User
	err := s.db.Model(&models.User{}).Where("id = ?", theID).First(&user).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error fetching user: %v", err)
		return []models.Downgrade{}, response.ErrorBadRequest(nil, "Could not find user")
	}

	var downgrades []models.Downgrade
	err = s.db.Model(&models.Downgrade{}).Where("user_id = ?", theID).Find(&downgrades).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error fetching downgrades: %v", err)
		return []models.Downgrade{}, response.ErrorBadRequest(nil, "Could not find downgrades")
	}

	return downgrades, response.ErrorResponse{} // Return success response
}
