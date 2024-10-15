package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

/*
 * @apiDefine: CreatePassiveRevenueCalculatorRequest
 */
type CreatePassiveRevenueCalculatorRequest struct {
	CompanyName     string `json:"company_name" openapi:"example:Company Name;type:string;"`
	FullName        string `json:"full_name" openapi:"example:James Mada;type:string;"`
	Email           string `json:"email" openapi:"example:JnY6A@example.com;type:string;"`
	LevelOneCount   int    `json:"level_one_count" openapi:"example:5;type:integer;"`
	LevelTwoCount   int    `json:"level_two_count" openapi:"example:5;type:integer;"`
	LevelThreeCount int    `json:"level_three_count" openapi:"example:5;type:integer;"`
	LevelFourCount  int    `json:"level_four_count" openapi:"example:5;type:integer;"`
	LevelFiveCount  int    `json:"level_five_count" openapi:"example:5;type:integer;"`
}

func (c *CreatePassiveRevenueCalculatorRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"company_name":      gocriteria.New("company_name").Required().MinMaxLength(2, 200),
		"full_name":         gocriteria.New("full_name").Required().MinMaxLength(2, 200),
		"email":             gocriteria.New("email").Required().Email(),
		"level_one_count":   gocriteria.New("level_one_count").Required(),
		"level_two_count":   gocriteria.New("level_two_count").Required(),
		"level_three_count": gocriteria.New("level_three_count").Required(),
		"level_four_count":  gocriteria.New("level_four_count").Required(),
		"level_five_count":  gocriteria.New("level_five_count").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"company_name":      "Company Name",
			"full_name":         "Full Name",
			"email":             "Email",
			"level_one_count":   "Level One Count",
			"level_two_count":   "Level Two Count",
			"level_three_count": "Level Three Count",
			"level_four_count":  "Level Four Count",
			"level_five_count":  "Level Five Count",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, c)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) CreatePassiveRevenue(ctx context.Context, input CreatePassiveRevenueCalculatorRequest) (models.PassiveRevenue, response.ErrorResponse) {

	passiveRevenue := models.PassiveRevenue{
		CompanyName:     input.CompanyName,
		FullName:        input.FullName,
		Email:           input.Email,
		LevelOneCount:   input.LevelOneCount,
		LevelTwoCount:   input.LevelTwoCount,
		LevelThreeCount: input.LevelThreeCount,
		LevelFourCount:  input.LevelFourCount,
		LevelFiveCount:  input.LevelFiveCount,
	}
	err := s.db.WithContext(ctx).Create(&passiveRevenue).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return passiveRevenue, response.GormErrorResponse(err, "An error occurred while creating the calculator")
	}
	return passiveRevenue, response.ErrorResponse{}
}
