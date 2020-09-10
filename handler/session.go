package handler

import (
	"elipzis.com/inertia-echo/util"
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

//
func (this *Handler) setSession(c echo.Context, key string, value interface{}, options *sessions.Options) error {
	// Register whatever object to be saved
	gob.Register(value)

	// Set it
	s, _ := session.Get("session", c)
	if options != nil {
		s.Options = options
	}
	s.Values[key] = value
	return s.Save(c.Request(), c.Response())
}

//
func (this *Handler) deleteSession(c echo.Context) error {
	s, _ := session.Get("session", c)
	s.Options.MaxAge = -1
	return s.Save(c.Request(), c.Response())
}

//
func (this *Handler) setErrors(c echo.Context, value *util.Error) error {
	s, _ := session.Get("session", c)
	s.AddFlash(value.ToString(), "_errors")
	return s.Save(c.Request(), c.Response())
}

//
func (this *Handler) addFlash(c echo.Context, value interface{}, vars ...string) error {
	s, _ := session.Get("session", c)
	s.AddFlash(value, vars...)
	return s.Save(c.Request(), c.Response())
}

//
func (this *Handler) addSuccessFlash(c echo.Context, value interface{}) *Handler {
	this.addFlash(c, value, "_flash_success")
	return this
}

//
func (this *Handler) addErrorFlash(c echo.Context, value interface{}) *Handler {
	this.addFlash(c, value, "_flash_error")
	return this
}

//
func (this *Handler) addWarningFlash(c echo.Context, value interface{}) *Handler {
	this.addFlash(c, value, "_flash_warning")
	return this
}
