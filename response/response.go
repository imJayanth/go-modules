package response

import (
	"github.com/imJayanth/go-modules/errors"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data          interface{}         `json:"data"`
	ResponseError errors.RestAPIError `json:"error"`
}

func NewResponse(data interface{}, error errors.RestAPIError) *Response {
	return &Response{Data: data, ResponseError: error}
}

func MapResponseFromRestAPIError(c *fiber.Ctx, restErr errors.RestAPIError) error {
	switch restErr.Status {
	case 401:
		return RespondUnAuthorised(c, restErr.Message)
	case 400:
		return RespondBadRequest(c, restErr.Message)
	case 404:
		return RespondNotFound(c, restErr.Message)
	case 403:
		return RespondForbidden(c, restErr.Message)
	case 409:
		return RespondDuplicate(c, restErr.Message)
	default:
		return RespondInternalServerError(c, restErr.Message)
	}
}
