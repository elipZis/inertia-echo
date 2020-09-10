package util

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"regexp"
	"strconv"
	"strings"
)

//
var keyErrorRegex, _ = regexp.Compile(`Key:(?P<Key>.*)Error:(?P<Value>.*)`)

// Error container for return messages
type Error struct {
	Errors map[string][]interface{} `json:"errors"`
}

// Create nicely formatted errors to return
func NewError() (this *Error) {
	this = new(Error)
	this.Errors = make(map[string][]interface{})
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
		break
	case validator.ValidationErrors:
		key, message = GetErrorsFromString(v.Error())
	}

	// this.Errors[key] = message
	if errors, ok := this.Errors[key]; ok {
		this.Errors[key] = append(errors, message)
	} else {
		this.Errors[key] = []interface{}{message}
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
	return c.Render(code, name, map[string]interface{}{
		"errors": this.Errors,
	})
}

// Return the errors map as key/error string
func (this *Error) ToString() []string {
	var errors []string
	for k, v := range this.Errors {
		for _, message := range v {
			errors = append(errors, fmt.Sprintf("Key: %s Error: %s", k, message))
		}
	}
	return errors
}

//
func GetErrorsFromString(err string) (string, string) {
	var key, value string
	res := keyErrorRegex.FindStringSubmatch(err)
	names := keyErrorRegex.SubexpNames()
	for i, _ := range res {
		if i != 0 {
			if names[i] == "Key" {
				key = strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(res[i]), "'"), "'")
			}
			if names[i] == "Value" {
				value = strings.TrimSpace(res[i])
			}
		}
	}
	return key, value
}
