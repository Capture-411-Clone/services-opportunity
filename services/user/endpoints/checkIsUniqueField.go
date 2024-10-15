package endpoints

import (
	"context"
	"errors"
	"fmt"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

/*
 * @apiDefine: GetCheckIsUniqueFieldParams
 */

type GetCheckIsUniqueFieldParams struct {
	Field string `json:"field" openapi:"example:field;nullable;in:path"`
	Value string `json:"value" openapi:"example:value;nullable;in:path"`
}

func (p *GetCheckIsUniqueFieldParams) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"field": gocriteria.New("field").Required(),
		"value": gocriteria.New("value").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"field": "Field",
			"value": "Value",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}
func (s *service) CheckIsUniqueField(ctx context.Context, field, value string) (
	bool, models.User, response.ErrorResponse,
) {
	var exists bool
	var user models.User
	var count int64

	// Validate the field against a list of allowed fields
	allowedFields := map[string]bool{
		"username": true,
		"email":    true,
		"phone":    true,
	}

	if !allowedFields[field] {
		return false, user, response.ErrorBadRequest(nil, "Invalid field")
	}

	err := s.db.WithContext(ctx).Where(fmt.Sprintf("%s = ?", field), value).
		First(&user).Count(&count).Error
	exists = count > 0
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.With(ctx).Error(err)
		return exists, user, response.GormErrorResponse(err, "An error occurred while finding the user")
	}
	return exists, user, response.ErrorResponse{}
}
