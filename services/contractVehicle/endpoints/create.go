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
 * @apiDefine: CreateContractVehicleRequest
 */
type CreateContractVehicleRequest struct {
	Name    string `json:"name" openapi:"example:Car;type:string;"`
	Acronym string `json:"acronym" openapi:"example:CV;type:string;"`
}

func (c *CreateContractVehicleRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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

func (s *service) Create(ctx context.Context, input CreateContractVehicleRequest) (models.ContractVehicle, response.ErrorResponse) {
	if !policy.CanCreateContractVehicle(ctx) {
		s.logger.With(ctx).Error("You do not have permission to create a contractVehicle")
		return models.ContractVehicle{}, response.ErrorForbidden("You do not have permission to create a contractVehicle")
	}

	Id := policy.ExtractIdClaim(ctx)
	id, _ := strconv.Atoi(Id)

	contractVehicle := models.ContractVehicle{
		UserID:  id,
		Name:    input.Name,
		Acronym: input.Acronym,
	}
	err := s.db.WithContext(ctx).Create(&contractVehicle).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return contractVehicle, response.GormErrorResponse(err, "An error occurred while creating the contractVehicle")
	}
	return contractVehicle, response.ErrorResponse{}
}
