package endpoints

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: CreateOfficeRequest
 */

type CreateOfficeRequest struct {
	ParentID int    `json:"parent_id" openapi:"example:1;type:integer"`
	AgencyID int    `json:"agency_id" openapi:"example:1;type:integer"`
	Acronym  string `json:"acronym" openapi:"example:TECH;type:string;"`
	Name     string `json:"name" openapi:"example:Technology Office;type:string"`
}

func (c *CreateOfficeRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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

func (s *service) Create(ctx context.Context, input CreateOfficeRequest) (models.Office, response.ErrorResponse) {
	if !policy.CanCreateOffice(ctx) {
		s.logger.With(ctx).Error("You do not have permission to create a office")
		return models.Office{}, response.ErrorForbidden("You do not have permission to create a office")
	}

	Id := policy.ExtractIdClaim(ctx)
	id, _ := strconv.Atoi(Id)

	office := models.Office{
		ParentID: dtp.NullInt64{
			Int64: int64(input.ParentID),
			Valid: input.ParentID != 0,
		},
		AgencyID: dtp.NullInt64{
			Int64: int64(input.AgencyID),
			Valid: input.AgencyID != 0,
		},
		UserID:  id,
		Acronym: input.Acronym,
		Name:    input.Name,
	}
	err := s.db.WithContext(ctx).Create(&office).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return office, response.GormErrorResponse(err, "An error occurred while creating the office")
	}
	return office, response.ErrorResponse{}
}
