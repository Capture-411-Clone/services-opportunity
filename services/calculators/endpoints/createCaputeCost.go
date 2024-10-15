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
 * @apiDefine: CreateCaptureCostCalculatorRequest
 */
type CreateCaptureCostCalculatorRequest struct {
	CompanyName           string `json:"company_name" openapi:"example:Company Name;type:string;"`
	FullName              string `json:"full_name" openapi:"example:James Mada;type:string;"`
	Email                 string `json:"email" openapi:"example:JnY6A@example.com;type:string;"`
	ManagerCount          int    `json:"manager_count" openapi:"example:5;type:integer;"`
	HourlyRate            int    `json:"hourly_rate" openapi:"example:100;type:integer;"`
	AnnualPipelineTime    int    `json:"annual_pipeline_time" openapi:"example:100;type:integer;"`
	CompanyAnnualCost     int    `json:"company_annual_cost" openapi:"example:100;type:integer;"`
	CompanyCapture411Cost int    `json:"company_capture411_cost" openapi:"example:100;type:integer;"`
	CompanySavings        int    `json:"company_savings" openapi:"example:100;type:integer;"`
	CompanyHoursSaved     int    `json:"company_hours_saved" openapi:"example:100;type:integer;"`
}

func (c *CreateCaptureCostCalculatorRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"company_name":            gocriteria.New("company_name").Required().MinMaxLength(2, 200),
		"full_name":               gocriteria.New("full_name").Required().MinMaxLength(2, 200),
		"email":                   gocriteria.New("email").Required().Email(),
		"manager_count":           gocriteria.New("manager_count").Required(),
		"hourly_rate":             gocriteria.New("hourly_rate").Required(),
		"annual_pipeline_time":    gocriteria.New("annual_pipeline_time").Required(),
		"company_annual_cost":     gocriteria.New("company_annual_cost").Required(),
		"company_capture411_cost": gocriteria.New("company_capture411_cost").Required(),
		"company_savings":         gocriteria.New("company_savings").Required(),
		"company_hours_saved":     gocriteria.New("company_hours_saved").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"company_name":            "Company Name",
			"full_name":               "Full Name",
			"email":                   "Email",
			"manager_count":           "Manager Count",
			"hourly_rate":             "Hourly Rate",
			"annual_pipeline_time":    "Annual Pipeline Time",
			"company_annual_cost":     "Company Annual Cost",
			"company_capture411_cost": "Company Capture411 Cost",
			"company_savings":         "Company Savings",
			"company_hours_saved":     "Company Hours Saved",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, c)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) CreateCaptureCost(ctx context.Context, input CreateCaptureCostCalculatorRequest) (models.CaptureCost, response.ErrorResponse) {

	captureCost := models.CaptureCost{
		CompanyName:           input.CompanyName,
		FullName:              input.FullName,
		Email:                 input.Email,
		ManagerCount:          input.ManagerCount,
		HourlyRate:            input.HourlyRate,
		AnnualPipelineTime:    input.AnnualPipelineTime,
		CompanyAnnualCost:     input.CompanyAnnualCost,
		CompanyCapture411Cost: input.CompanyCapture411Cost,
		CompanySavings:        input.CompanySavings,
		CompanyHoursSaved:     input.CompanyHoursSaved,
	}
	err := s.db.WithContext(ctx).Create(&captureCost).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return captureCost, response.GormErrorResponse(err, "An error occurred while creating the calculator")
	}
	return captureCost, response.ErrorResponse{}
}
