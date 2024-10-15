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
 * @apiDefine: UpdateAgencyRequest
 */
type UpdateAgencyRequest struct {
	Name         string `json:"name" openapi:"example:My Agency;type:string;"`
	Acronym      string `json:"acronym" openapi:"example:TDE;type:string"`
	DepartmentID int    `json:"department_id" openapi:"example:1;type:int"`
}

func (c *UpdateAgencyRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"name":          gocriteria.New("name").Required().MinMaxLength(2, 200),
		"acronym":       gocriteria.New("acronym").Optional(),
		"department_id": gocriteria.New("department_id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"name":          "Name",
			"acronym":       "Acronym",
			"Department_id": "Department",
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
 * @apiDefine: AgencyRequestUpdateParams
 */
type AgencyRequestUpdateParams struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *AgencyRequestUpdateParams) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
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

func (s *service) Update(ctx context.Context, id string, input UpdateAgencyRequest) (
	models.Agency, response.ErrorResponse,
) {
	var agency models.Agency

	if !policy.CanUpdateAgency(ctx) {
		s.logger.With(ctx).Error("You do not have permission to update a agency")
		return models.Agency{}, response.ErrorForbidden(nil, "You do not have permission to update a agency")
	}

	err := s.db.WithContext(ctx).First(&agency, "id", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Agency{}, response.GormErrorResponse(err, "An error occurred while finding the agency")
	}
	agency.Name = input.Name
	agency.Acronym = input.Acronym
	agency.DepartmentID = input.DepartmentID

	err = s.db.WithContext(ctx).Save(&agency).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Agency{}, response.GormErrorResponse(err, "An error occurred while updating the agency")
	}
	return agency, response.ErrorResponse{}
}
