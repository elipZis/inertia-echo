package util

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"strings"
)

//
var DefaultSessionOptions = &sessions.Options{
	MaxAge: 86400 * 7,
}

//
func SetSession(c echo.Context, key string, value interface{}, options *sessions.Options) error {
	// Register whatever object to be saved
	gob.Register(value)

	// Set it
	sess, _ := session.Get("session", c)
	if options == nil {
		options = DefaultSessionOptions
	}
	sess.Options = options
	sess.Values[key] = value
	return sess.Save(c.Request(), c.Response())
}

//
func DeleteSession(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	return sess.Save(c.Request(), c.Response())
}

//
func GetRedirectUrl(c echo.Context, path string) string {
	return GetBaseUrl(c) + path
}

//
func GetBaseUrl(c echo.Context) string {
	req, scheme := c.Request(), c.Scheme()
	host := req.Host
	url := scheme + "://" + host // + req.RequestURI
	return strings.TrimSuffix(url, "/")
}
