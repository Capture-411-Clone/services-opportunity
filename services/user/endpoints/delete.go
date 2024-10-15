package endpoints

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: DeleteUserRequest
 */
type DeleteUserRequest struct {
	IDs string `json:"ids" openapi:"example:1,2,3;"`
}

func (a *DeleteUserRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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
func (s *service) Delete(ctx context.Context, input DeleteUserRequest) (string, response.ErrorResponse) {
	Id := policy.ExtractIdClaim(ctx)
	var adminUser models.User
	err := s.db.WithContext(ctx).First(&adminUser, "id", Id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in the database")
	}

	if config.AppConfig.Environment != "development" {
		s.logger.With(ctx).Error("you are not in development environment")
		return "", response.ErrorForbidden("You do not have permission to access this user")
	}

	if !policy.CanDeleteUser(ctx) {
		return "", response.ErrorForbidden("You do not have permission to access this user")
	}

	var ids []int
	for _, idStr := range strings.Split(input.IDs, ",") {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorResponse{
				Message: "Invalid input IDs",
			}
		}
		ids = append(ids, id)
	}

	// find first
	var users []models.User
	err = s.db.WithContext(ctx).Where("id IN ?", ids).Find(&users).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete user")
	}
	if len(users) == 0 {
		return "", response.ErrorResponse{
			Message: "Not found any user",
		}
	}

	err = s.db.WithContext(ctx).Where("id IN ?", ids).Delete(&models.User{}).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete user")
	}

	return input.IDs, response.ErrorResponse{}
}
