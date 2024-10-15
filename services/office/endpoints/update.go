package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: UpdateOfficeRequest
 */
type UpdateOfficeRequest struct {
	ParentID int    `json:"parent_id" openapi:"example:1;type:integer"`
	AgencyID int    `json:"agency_id" openapi:"example:1;type:integer"`
	Acronym  string `json:"acronym" openapi:"example:TECH;type:string"`
	Name     string `json:"name" openapi:"example:Technology Office;type:string"`
}

func (c *UpdateOfficeRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"parent_id": gocriteria.New("parent_id").Optional(),
		"agency_id": gocriteria.New("agency_id").Optional(),
		"acronym":   gocriteria.New("acronym").Optional(),
		"name":      gocriteria.New("name").Required().MinMaxLength(2, 200),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"parent_id": "Related Office",
			"agency_id": "Agency",
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
 * @apiDefine: UpdateOfficeRequestParam
 */
type UpdateOfficeRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *UpdateOfficeRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
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

func (s *service) Update(ctx context.Context, id string, input UpdateOfficeRequest) (
	models.Office, response.ErrorResponse,
) {
	var office models.Office

	if !policy.CanUpdateOffice(ctx) {
		s.logger.With(ctx).Error("You do not have permission to update a office")
		return models.Office{}, response.ErrorForbidden(nil, "You do not have permission to update a office")
	}

	err := s.db.WithContext(ctx).First(&office, "id", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Office{}, response.GormErrorResponse(err, "An error occurred while finding the office")
	}
	office.ParentID = dtp.NullInt64{
		Int64: int64(input.ParentID),
		Valid: input.ParentID != 0,
	}
	office.AgencyID = dtp.NullInt64{
		Int64: int64(input.AgencyID),
		Valid: input.AgencyID != 0,
	}
	office.Acronym = input.Acronym
	office.Name = input.Name

	err = s.db.WithContext(ctx).Save(&office).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Office{}, response.GormErrorResponse(err, "An error occurred while updating the office")
	}
	return office, response.ErrorResponse{}
}
