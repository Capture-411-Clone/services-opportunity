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
 * @apiDefine: DeleteDepartmentRequest
 */
type DeleteDepartmentRequest struct {
	IDs string `json:"ids" openapi:"example:1,2,3;"`
}

func (a *DeleteDepartmentRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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

func (s *service) Delete(ctx context.Context, input DeleteDepartmentRequest) (string, response.ErrorResponse) {
	if !policy.CanDeleteDepartment(ctx) {
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
	var departments []models.Department
	err := s.db.WithContext(ctx).Where("id IN ?", ids).Find(&departments).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete department")
	}
	if len(departments) == 0 {
		return "", response.ErrorResponse{
			Message: "Not found any department",
		}
	}

	err = s.db.WithContext(ctx).Where("id IN ?", ids).Delete(&models.Department{}).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete department")
	}

	return input.IDs, response.ErrorResponse{}
}
