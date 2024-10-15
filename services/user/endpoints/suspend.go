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
 * @apiDefine: SuspendRequestParam
 */
type SuspendRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *SuspendRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
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
func (s *service) Suspend(ctx context.Context, id string) (models.User, response.ErrorResponse) {
	var user models.User
	if !policy.CanSuspendUser(ctx) {
		s.logger.With(ctx).Error("You do not have permission to access this user")
		err := response.ErrorForbidden("You do not have permission to access this user")
		return user, err
	}

	err := s.db.WithContext(ctx).First(&user, "id", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return user, response.GormErrorResponse(err, "Error in finding the user")
	}

	if user.SuspendedAt.Valid {
		user.SuspendedAt = dtp.NullTime{}
	} else {
		user.SuspendedAt = faker.SQLNow()
	}

	err = s.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return user, response.GormErrorResponse(err, "Error in saving the user")
	}

	return user, response.ErrorResponse{}
}
