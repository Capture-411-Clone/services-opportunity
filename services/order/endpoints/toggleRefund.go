package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/kit/faker"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: ToggleRefundRequestParam
 */
type ToggleRefundRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *ToggleRefundRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
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

func (s *service) ToggleRefund(ctx context.Context, id string) (models.Order, response.ErrorResponse) {
	if !policy.CanToggleRefundAt(ctx) {
		s.logger.With(ctx).Error("You do not have permission to change the refund at")
		return models.Order{}, response.ErrorForbidden("You do not have permission to change the refund at")
	}

	order := models.Order{}

	err := s.db.WithContext(ctx).First(&order, "id", id).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return order, response.GormErrorResponse(err, "Error in finding the order")
	}

	if order.RefundedAt.Valid {
		order.RefundedAt = dtp.NullTime{}
	} else {
		order.RefundedAt = faker.SQLNow()
	}

	err = s.db.WithContext(ctx).Save(&order).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return order, response.GormErrorResponse(err, "Error in saving the order")
	}

	return order, response.ErrorResponse{}

}
