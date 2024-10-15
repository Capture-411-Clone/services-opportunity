package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
)

/*
 * @apiDefine: ApproveEmailRequestParam
 */
type ApproveEmailRequestParam struct {
	Token string `json:"token" openapi:"example:3e2246ee4315c7e3a6032;nullable;in:path"`
}

func (p *ApproveEmailRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"token": gocriteria.New("token").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"token": "Token",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}
func (s *service) ApproveEmail(ctx context.Context, token string) (string, response.ErrorResponse) {
	panic("implement me")
}
