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
 * @apiDefine: CreateEntityHuntRequest
 */
type CreateEntityHuntRequest struct {
	Details            models.JSONMap `json:"details" openapi:"$ref:JSONMap;example='{\"name\": \"rey\", \"age\": 30, \"email\": \"rey@example.com\"}';type=object"`
	SolicitationNumber string         `json:"solicitation_number" openapi:"example:123456;type:string;"`
	Year               string         `json:"year" gorm:"size:4;uniqueIndex:udx_entity_hunts" openapi:"example:2022;type:string;"`
	FileName           string         `json:"file_name" openapi:"example:filename;type:string;"`
	FilePath           string         `json:"file_path" openapi:"example:user/documents/report.pdf;type:string"`
}

func (c *CreateEntityHuntRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"details":             gocriteria.New("details").Required(),
		"solicitation_number": gocriteria.New("solicitation_number").Optional(),
		"year":                gocriteria.New("year").Optional(),
		"file_name":           gocriteria.New("file_name").Optional(),
		"file_path":           gocriteria.New("file_path").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"details":             "Details",
			"solicitation_number": "Solicitation Number",
			"year":                "Year",
			"file_name":           "File Name",
			"file_path":           "File Path",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, c)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Create(ctx context.Context, input CreateEntityHuntRequest) (models.EntityHunt, response.ErrorResponse) {
	if !policy.CanCreateEntityHunt(ctx) {
		s.logger.With(ctx).Error("You do not have permission to create a entityHunt")
		return models.EntityHunt{}, response.ErrorForbidden("You do not have permission to create a entityHunt")
	}

	var entityHunt = models.EntityHunt{
		Details:            input.Details,
		SolicitationNumber: input.SolicitationNumber,
		Year:               input.Year,
		FileName:           input.FileName,
		FilePath:           input.FilePath,
	}
	err := s.db.WithContext(ctx).Create(&entityHunt).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return entityHunt, response.GormErrorResponse(err, "An error occurred while creating the entityHunt")
	}
	return entityHunt, response.ErrorResponse{}
}
