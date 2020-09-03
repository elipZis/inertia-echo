package inertia

import (
	"github.com/elipzis/inertia-echo/util"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

//
type MiddlewareConfig struct {
	Inertia *Inertia
	Skipper middleware.Skipper
}

// The Inertia Middleware to check every request for what it needs
func MiddlewareWithConfig(config MiddlewareConfig) echo.MiddlewareFunc {
	if config.Inertia == nil {
		config.Inertia = NewInertia()
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip, if configured and true
			if config.Skipper != nil && config.Skipper(c) {
				return next(c)
			}

			// Run Inertia post
			if err := next(c); err != nil {
				c.Error(err)
				e := util.NewError().AddError(err)
				// Add general errors in case some pop up
				config.Inertia.Share("errors", e.Errors)
			}

			req := c.Request()
			res := c.Response()

			if req.Header.Get(HeaderPrefix) == "" {
				return nil
			}

			if req.Method == "GET" && req.Header.Get(HeaderVersion) != config.Inertia.GetVersion() {
				// $request->session()->reflash();???
				if s, err := session.Get("session", c); err == nil {
					flashes := s.Flashes()
					config.Inertia.Share("flash", flashes)
					for _, flash := range flashes {
						s.AddFlash(flash)
					}
				}

				res.Header().Set(HeaderLocation, req.URL.String())
				return c.String(http.StatusConflict, "")
			}

			if exists, _ := util.InArray(req.Method, []string{"PUT", "PATCH", "DELETE"}); exists && res.Status == 302 {
				res.Status = http.StatusSeeOther
			}

			return nil
		}
	}
}
