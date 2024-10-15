package endpoints

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: CreateNaicsRequest
 */
type CreateNaicsRequest struct {
	Name string `json:"name" openapi:"example:Technology Department;type:string"`
}

func (c *CreateNaicsRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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

func (s *service) Create(ctx context.Context, input CreateNaicsRequest) (models.Naics, response.ErrorResponse) {
	if !policy.CanCreateNaics(ctx) {
		s.logger.With(ctx).Error("You do not have permission to create a naics")
		return models.Naics{}, response.ErrorForbidden("You do not have permission to create a naics")
	}

	Id := policy.ExtractIdClaim(ctx)
	id, _ := strconv.Atoi(Id)

	naics := models.Naics{
		UserID: id,
		Name:   input.Name,
	}
	err := s.db.WithContext(ctx).Create(&naics).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return naics, response.GormErrorResponse(err, "An error occurred while creating the naics")
	}
	return naics, response.ErrorResponse{}
}
