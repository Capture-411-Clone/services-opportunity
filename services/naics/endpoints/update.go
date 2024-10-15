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
 * @apiDefine: UpdateNaicsRequest
 */
type UpdateNaicsRequest struct {
	Name string `json:"name" openapi:"example:Technology Department;type:string"`
}

func (c *UpdateNaicsRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"name": gocriteria.New("name").Required().MinMaxLength(2, 200),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"name": "Name",
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
 * @apiDefine: UpdateNaicsRequestParam
 */
type UpdateNaicsRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *UpdateNaicsRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
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

func (s *service) Update(ctx context.Context, id string, input UpdateNaicsRequest) (
	models.Naics, response.ErrorResponse,
) {
	var naics models.Naics

	if !policy.CanUpdateNaics(ctx) {
		s.logger.With(ctx).Error("You do not have permission to update a naics")
		return models.Naics{}, response.ErrorForbidden(nil, "You do not have permission to update a naics")
	}

	err := s.db.WithContext(ctx).First(&naics, "id", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Naics{}, response.GormErrorResponse(err, "An error occurred while finding the naics")
	}
	naics.Name = input.Name

	err = s.db.WithContext(ctx).Save(&naics).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Naics{}, response.GormErrorResponse(err, "An error occurred while updating the naics")
	}
	return naics, response.ErrorResponse{}
}
