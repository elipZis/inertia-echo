package middleware

import (
	"elipzis.com/inertia-echo/util"
	"github.com/elipzis/inertia-echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"reflect"
)

//
type SessionMiddlewareConfig struct {
	Inertia *inertia.Inertia
	Skipper middleware.Skipper
}

//
func SessionMiddleware(inertia *inertia.Inertia) echo.MiddlewareFunc {
	return SessionMiddlewareWithConfig(SessionMiddlewareConfig{
		Inertia: inertia,
	})
}

//
func SessionMiddlewareWithConfig(config SessionMiddlewareConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip, if configured and true
			if config.Skipper != nil && config.Skipper(c) {
				return next(c)
			}

			// Flash a-ah
			shareFlashes(c, config.Inertia)
			shareErrors(c, config.Inertia)

			// return next(c)

			// Run
			if err := next(c); err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()
			if req.Method == "GET" && req.Header.Get(inertia.HeaderVersion) != config.Inertia.GetVersion() {
				// Reflash?
				if s, err := session.Get("session", c); err == nil {
					flashes := s.Flashes()
					_ = s.Save(c.Request(), c.Response())
					// config.Inertia.Share("flash", flashes)
					for _, flash := range flashes {
						s.AddFlash(flash)
					}
				}

				res.Header().Set(inertia.HeaderLocation, req.URL.String())
				return c.String(http.StatusConflict, "")
			}

			return nil
		}
	}
}

// Cleans and adds any new session flashes
func shareFlashes(c echo.Context, inertia *inertia.Inertia) {
	inertia.Share(c, "flash", map[string]interface{}{})
	if s, err := session.Get("session", c); err == nil {
		inertia.Share(c, "flash", map[string]interface{}{
			"success": s.Flashes("_flash_success"),
			"error":   s.Flashes("_flash_error"),
			"warning": s.Flashes("_flash_warning"),
		})
		_ = s.Save(c.Request(), c.Response())
	}
}

// Cleans and adds any new session errors
func shareErrors(c echo.Context, inertia *inertia.Inertia) {
	inertia.Share(c, "errors", map[string]interface{}{})
	if s, err := session.Get("session", c); err == nil {
		errorFlashes := s.Flashes("_errors")
		_ = s.Save(c.Request(), c.Response())
		if errorFlashes != nil {
			switch reflect.TypeOf(errorFlashes).Kind() {
			case reflect.Slice:
				if len(errorFlashes) > 0 {
					errors := make(map[string]interface{})
					flashes := errorFlashes[0].([]string)
					for _, v := range flashes {
						key, value := util.GetErrorsFromString(v)
						if val, ok := errors[key]; ok {
							errors[key] = append(val.([]string), value)
						} else {
							errors[key] = []string{value}
						}
					}
					inertia.Share(c, "errors", errors)
				}
			}
		}
	}
}
