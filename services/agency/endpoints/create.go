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
 * @apiDefine: CreateAgencyRequest
 */
type CreateAgencyRequest struct {
	Name         string `json:"name" openapi:"example:My Agency;type:string;"`
	Acronym      string `json:"acronym" openapi:"example:TDE;type:string"`
	DepartmentID int    `json:"department_id" openapi:"example:1;type:int"`
}

func (c *CreateAgencyRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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

func (s *service) Create(ctx context.Context, input CreateAgencyRequest) (models.Agency, response.ErrorResponse) {
	if !policy.CanCreateAgency(ctx) {
		s.logger.With(ctx).Error("You do not have permission to create a agency")
		return models.Agency{}, response.ErrorForbidden("You do not have permission to create a agency")
	}

	Id := policy.ExtractIdClaim(ctx)
	id, _ := strconv.Atoi(Id)

	agency := models.Agency{
		UserID:       id,
		Name:         input.Name,
		Acronym:      input.Acronym,
		DepartmentID: input.DepartmentID,
	}
	err := s.db.WithContext(ctx).Create(&agency).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return agency, response.GormErrorResponse(err, "An error occurred while creating the agency")
	}
	return agency, response.ErrorResponse{}
}
