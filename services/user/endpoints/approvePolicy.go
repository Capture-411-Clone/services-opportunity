package endpoints

import (
	"context"
	"strconv"
	"time"

	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

func (s *service) ApprovePolicy(ctx context.Context) (models.User, response.ErrorResponse) {
	var user models.User
	tx := s.db.WithContext(ctx).Begin() // Begin a new transaction
	if tx.Error != nil {
		return models.User{}, response.GormErrorResponse(tx.Error, "An error occurred on the server")
	}

	Id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(Id)

	err := tx.Preload("Roles").First(&user, "id", theID).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.User{}, response.GormErrorResponse(err, "Error in finding the user")
	}

	user.PolicyApprovedAt = dtp.NullTime{Valid: true, Time: time.Now()}

	err = tx.Save(&user).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.User{}, response.GormErrorResponse(err, "Error in saving the user")
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback() // Rollback the transaction if there's any error or panic
		} else {
			err = tx.Commit().Error // Commit the transaction if there are no errors
		}
	}()
	return user, response.ErrorResponse{}
}
