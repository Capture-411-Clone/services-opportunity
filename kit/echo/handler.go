package echo

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/labstack/echo/v4"
)

func CustomHandler(err error, c echo.Context) {
	r := buildErrorResponse(err)

	if err := c.JSON(r.StatusCode, err); err != nil {
		c.Logger().Errorf("failed writing error response: %v", err)
	}
}

// buildErrorResponse builds an error response from an error.
func buildErrorResponse(err error) response.ErrorResponse {
	switch err := err.(type) {
	case *echo.HTTPError:
		switch err.Code {
		case http.StatusNotFound:
			return response.ErrorNotFound(nil, "The requested resource was not found")
		default:
			return response.ErrorResponse{
				StatusCode: err.Code,
				Data:       nil,
				Message:    "An error occurred while processing the request",
			}
		}
	}

	return response.ErrorInternalServerError(nil, "Error occurred while processing the request")
}
