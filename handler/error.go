package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Error container for return messages
type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

//
func (this *Handler) createErrorResponse(c echo.Context, err error, status ...int) error {
	code := http.StatusInternalServerError
	if len(status) > 0 {
		code = status[0]
	}
	return c.JSON(code, this.createError(err))
}

// Create a nicely formatted error to return
func (this *Handler) createError(err error) *Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}
	return &e
}

//
func AccessForbidden() *Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "access forbidden"
	return &e
}

//
func NotFound() *Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "resource not found"
	return &e
}
