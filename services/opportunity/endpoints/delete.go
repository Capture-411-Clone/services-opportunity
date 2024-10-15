package endpoints

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: DeleteOpportunityRequest
 */
type DeleteOpportunityRequest struct {
	IDs string `json:"ids" openapi:"example:1,3,4"`
}

func (a *DeleteOpportunityRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"ids": gocriteria.New("ids").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"ids": "IDs",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, a)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Delete(ctx context.Context, input DeleteOpportunityRequest) (string, response.ErrorResponse) {

	// Parse and validate IDs
	var ids []int
	for _, idStr := range strings.Split(input.IDs, ",") {
		id, err := strconv.Atoi(strings.TrimSpace(idStr))
		if err != nil {
			s.logger.With(ctx).Errorf("Invalid ID: %v", err)
			return "", response.ErrorResponse{
				Message: "Invalid input IDs",
			}
		}
		ids = append(ids, id)
	}

	// Start a transaction
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		s.logger.With(ctx).Errorf("Error starting transaction: %v", tx.Error)
		return "", response.GormErrorResponse(tx.Error, "An error occurred while starting the transaction")
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			s.logger.With(ctx).Errorf("Transaction rolled back due to panic: %v", r)
		}
	}()

	// Retrieve opportunities
	var opportunities []models.Opportunity
	err := tx.Where("id IN ?", ids).Find(&opportunities).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error finding opportunities: %v", err)
		tx.Rollback()
		return "", response.GormErrorResponse(err, "An error occurred while retrieving opportunities")
	}
	if len(opportunities) == 0 {
		tx.Rollback()
		return "", response.ErrorResponse{
			Message: "No opportunities found",
		}
	}

	for _, opportunity := range opportunities {
		if !policy.CanDeleteOpportunity(ctx, opportunity) {
			s.logger.With(ctx).Error("You do not have permission to delete an opportunity")
			return "", response.ErrorForbidden("You do not have permission to delete opportunity " + opportunity.SolicitationNumber)
		}
	}

	// Update duplicated flag
	for _, opportunity := range opportunities {
		var remainingOpportunities []models.Opportunity
		err = tx.Model(&models.Opportunity{}).Where("solicitation_number = ? AND id NOT IN ?", opportunity.SolicitationNumber, ids).Find(&remainingOpportunities).Error
		if err != nil {
			s.logger.With(ctx).Errorf("Error finding remaining opportunities: %v", err)
			tx.Rollback()
			return "", response.GormErrorResponse(err, "An error occurred while checking remaining opportunities")
		}
		if len(remainingOpportunities) == 1 {
			err = tx.Model(&models.Opportunity{}).Where("id = ?", remainingOpportunities[0].ID).Update("duplicated", false).Error
			if err != nil {
				s.logger.With(ctx).Errorf("Error updating duplicated flag: %v", err)
				tx.Rollback()
				return "", response.GormErrorResponse(err, "An error occurred while updating the duplicated flag")
			}
		}
	}

	// Delete opportunities
	err = tx.Where("id IN ?", ids).Delete(&models.Opportunity{}).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error deleting opportunities: %v", err)
		tx.Rollback()
		return "", response.GormErrorResponse(err, "An error occurred while deleting opportunities")
	}

	// Commit transaction
	err = tx.Commit().Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error committing transaction: %v", err)
		return "", response.GormErrorResponse(err, "An error occurred while committing the transaction")
	}

	return input.IDs, response.ErrorResponse{}
}
