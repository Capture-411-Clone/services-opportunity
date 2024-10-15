package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

/*
 * @apiDefine: CheckDuplicateOpportunityRequestParam
 */
type CheckDuplicateOpportunityRequestParam struct {
	SolicitationNumber string `json:"solicitation_number" openapi:"example:123456;type:string;"`
}

func (p *CheckDuplicateOpportunityRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"solicitation_number": gocriteria.New("solicitation_number").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"solicitation_number": "Solicitation Number",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) CheckDuplicate(ctx context.Context, solicitation_number string) (
	bool, models.Opportunity, response.ErrorResponse,
) {

	var opportunitiesCount int64
	var opportunity models.Opportunity

	err := s.db.Model(&opportunity).Where("solicitation_number = ?", solicitation_number).Count(&opportunitiesCount).First(&opportunity).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error checking duplicate opportunity: %v", err)
		return true, opportunity, response.GormErrorResponse(err, "error checking duplicate opportunity")
	}

	if opportunitiesCount > 0 {
		return true, opportunity, response.ErrorResponse{}
	}

	return false, opportunity, response.ErrorResponse{}
}
