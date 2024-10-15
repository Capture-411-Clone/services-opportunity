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
 * @apiDefine: UpdateContractVehicleRequest
 */
type UpdateContractVehicleRequest struct {
	Name    string `json:"name" openapi:"example:Car;type:string;"`
	Acronym string `json:"acronym" openapi:"example:CV;type:string;"`
}

func (c *UpdateContractVehicleRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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
 * @apiDefine: ContractVehicleUpdateParams
 */
type ContractVehicleUpdateParams struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *ContractVehicleUpdateParams) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
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

func (s *service) Update(ctx context.Context, id string, input UpdateContractVehicleRequest) (
	models.ContractVehicle, response.ErrorResponse,
) {
	var contractVehicle models.ContractVehicle

	if !policy.CanUpdateContractVehicle(ctx) {
		s.logger.With(ctx).Error("You do not have permission to update a contractVehicle")
		return models.ContractVehicle{}, response.ErrorForbidden(nil, "You do not have permission to update a contractVehicle")
	}

	err := s.db.WithContext(ctx).First(&contractVehicle, "id", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.ContractVehicle{}, response.GormErrorResponse(err, "An error occurred while finding the contractVehicle")
	}
	contractVehicle.Name = input.Name
	contractVehicle.Acronym = input.Acronym

	err = s.db.WithContext(ctx).Save(&contractVehicle).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.ContractVehicle{}, response.GormErrorResponse(err, "An error occurred while updating the contractVehicle")
	}
	return contractVehicle, response.ErrorResponse{}
}
