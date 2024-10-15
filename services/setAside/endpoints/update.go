package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: UpdateSetAsideRequest
 */
type UpdateSetAsideRequest struct {
	Name    string `json:"name" openapi:"example:Technology;type:string"`
	Acronym string `json:"acronym" openapi:"example:TDE;type:string"`
}

func (c *UpdateSetAsideRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"name":    gocriteria.New("name").Required().MinMaxLength(2, 200),
		"acronym": gocriteria.New("acronym").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"name":    "Name",
			"acronym": "Acronym",
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
 * @apiDefine: UpdateSetAsideRequestParam
 */
type UpdateSetAsideRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *UpdateSetAsideRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
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

func (s *service) Update(ctx context.Context, id string, input UpdateSetAsideRequest) (
	models.SetAside, response.ErrorResponse,
) {
	var setAside models.SetAside

	if !policy.CanUpdateSetAside(ctx) {
		s.logger.With(ctx).Error("You do not have permission to update a setAside")
		return models.SetAside{}, response.ErrorForbidden(nil, "You do not have permission to update a setAside")
	}

	fmt.Println("setAsideid ", id)
	err := s.db.WithContext(ctx).First(&setAside, "id", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.SetAside{}, response.GormErrorResponse(err, "An error occurred while finding the setAside")
	}
	fmt.Println("setAsidefound ")
	setAside.Name = input.Name
	setAside.Acronym = input.Acronym

	err = s.db.WithContext(ctx).Save(&setAside).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.SetAside{}, response.GormErrorResponse(err, "An error occurred while updating the setAside")
	}
	return setAside, response.ErrorResponse{}
}
