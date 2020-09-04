package util

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

// Error container for return messages
type Error struct {
	Errors map[string][]string `json:"errors"`
}

// Create nicely formatted errors to return
func NewError() (this *Error) {
	this = new(Error)
	this.Errors = make(map[string][]string)
	return this
}

// Add an error
func (this *Error) AddError(err error, name ...string) *Error {
	key := "general"
	message := err.Error()
	if len(name) > 0 {
		key = name[0]
	}
	switch v := err.(type) {
	case *echo.HTTPError:
		if len(name) <= 0 {
			key = strconv.Itoa(v.Code)
		}
		message = v.Message.(string)
	}

	if errors, ok := this.Errors[key]; ok {
		this.Errors[key] = append(errors, message)
	} else {
		this.Errors[key] = []string{message}
	}

	return this
}

//
func (this *Error) HasErrors() bool {
	return len(this.Errors) > 0
}

// Write a json response of these errors to the given context
func (this *Error) JSON(c echo.Context, status ...int) error {
	code := 500
	if len(status) > 0 {
		code = status[0]
	}
	return c.JSON(code, this.Errors)
}

// Render a response of these errors to the given context by e.g. template name
func (this *Error) Render(c echo.Context, name string, status ...int) error {
	code := 500
	if len(status) > 0 {
		code = status[0]
	}
	return c.Render(code, name, this.Errors)
}
