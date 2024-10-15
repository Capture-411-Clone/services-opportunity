package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/pagination"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/kit/restypes"
	"github.com/Capture-411/services-opportunity/services/user/endpoints"
	"github.com/labstack/echo/v4"
)

func (r resource) query(ctx echo.Context) error {
	var params = &endpoints.UserQueryRequestParams{}

	errs := params.ValidateQueries(ctx.Request())
	if len(errs) > 0 {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errs, "داده های ورودی نامعتبر است"))
	}

	if params.OrderBy == "" {
		params.OrderBy = "id"
	}

	if params.Order == "" {
		params.Order = "desc"
	}

	count, err := r.service.Count(ctx.Request().Context(), *params)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	pages := pagination.NewFromRequest(ctx.Request(), int(count))
	users, er := r.service.Query(
		ctx.Request().Context(), pages.Offset(), pages.Limit(), *params)
	if er.StatusCode != 0 {
		return ctx.JSON(er.StatusCode, er)
	}

	result := restypes.QueryResponse{
		Limit:      pages.Limit(),
		Offset:     pages.Offset(),
		Page:       pages.Page,
		TotalRows:  int64(pages.TotalCount),
		TotalPages: pages.PageCount,
		Items:      users,
	}

	return ctx.JSON(http.StatusOK, response.Success(result))
}
