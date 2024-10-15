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
 * @apiDefine: CreateDepartmentRequest
 */
type CreateDepartmentRequest struct {
	MarketID int    `json:"market_id" openapi:"example:1;type:int"`
	Acronym  string `json:"acronym" openapi:"example:TDE;type:string"`
	Name     string `json:"name" openapi:"example:Technology Department;type:string"`
}

func (c *CreateDepartmentRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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

func (s *service) Create(ctx context.Context, input CreateDepartmentRequest) (models.Department, response.ErrorResponse) {
	if !policy.CanCreateDepartment(ctx) {
		s.logger.With(ctx).Error("You do not have permission to create a department")
		return models.Department{}, response.ErrorForbidden("You do not have permission to create a department")
	}

	Id := policy.ExtractIdClaim(ctx)
	id, _ := strconv.Atoi(Id)

	department := models.Department{
		UserID:   id,
		MarketID: input.MarketID,
		Acronym:  input.Acronym,
		Name:     input.Name,
	}
	err := s.db.WithContext(ctx).Create(&department).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return department, response.GormErrorResponse(err, "An error occurred while creating the department")
	}
	return department, response.ErrorResponse{}
}
