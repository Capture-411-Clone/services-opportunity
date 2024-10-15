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
 * @apiDefine: DeleteSetAsideRequest
 */
type DeleteSetAsideRequest struct {
	IDs string `json:"ids" openapi:"example:1,2,3;"`
}

func (a *DeleteSetAsideRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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

func (s *service) Delete(ctx context.Context, input DeleteSetAsideRequest) (string, response.ErrorResponse) {
	if !policy.CanDeleteSetAside(ctx) {
		s.logger.With(ctx).Error("You do not have permission to delete a department")
		return "", response.ErrorForbidden("You do not have permission to delete a department")
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
	var setAside []models.SetAside
	err := s.db.WithContext(ctx).Where("id IN ?", ids).Find(&setAside).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete set aside")
	}
	if len(setAside) == 0 {
		return "", response.ErrorResponse{
			Message: "Not found any set aside",
		}
	}

	err = s.db.WithContext(ctx).Where("id IN ?", ids).Delete(&models.SetAside{}).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete set aside")
	}

	return input.IDs, response.ErrorResponse{}
}
