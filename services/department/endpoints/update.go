package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: UpdateDepartmentRequest
 */
type UpdateDepartmentRequest struct {
	MarketID int    `json:"market_id" openapi:"example:1;type:int"`
	Acronym  string `json:"acronym" openapi:"example:TDE;type:string"`
	Name     string `json:"name" openapi:"example:Technology Department;type:string"`
}

func (c *UpdateDepartmentRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"market_id": gocriteria.New("market_id").Required(),
		"acronym":   gocriteria.New("acronym").Optional(),
		"name":      gocriteria.New("name").Required().MinMaxLength(2, 200),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"Market_id": "Market",
			"acronym":   "Acronym",
			"name":      "Name",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, c)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

/*
 * @apiDefine: UpdateDepartmentRequestParam
 */
type UpdateDepartmentRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *UpdateDepartmentRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id": gocriteria.New("id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"id": "ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Update(ctx context.Context, id string, input UpdateDepartmentRequest) (
	models.Department, response.ErrorResponse,
) {
	var department models.Department

	if !policy.CanUpdateDepartment(ctx) {
		s.logger.With(ctx).Error("You do not have permission to update a department")
		return models.Department{}, response.ErrorForbidden(nil, "You do not have permission to update a department")
	}

	err := s.db.WithContext(ctx).First(&department, "id", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Department{}, response.GormErrorResponse(err, "An error occurred while finding the department")
	}
	department.MarketID = input.MarketID
	department.Acronym = input.Acronym
	department.Name = input.Name

	err = s.db.WithContext(ctx).Save(&department).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Department{}, response.GormErrorResponse(err, "An error occurred while updating the department")
	}
	return department, response.ErrorResponse{}
}
