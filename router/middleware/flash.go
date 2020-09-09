package middleware

import (
	"elipzis.com/inertia-echo/inertia"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//
type FlashMiddlewareConfig struct {
	Inertia *inertia.Inertia
	Skipper middleware.Skipper
}

//
func FlashMiddleware(inertia *inertia.Inertia) echo.MiddlewareFunc {
	return FlashMiddlewareWithConfig(FlashMiddlewareConfig{
		Inertia: inertia,
	})
}

//
func FlashMiddlewareWithConfig(config FlashMiddlewareConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip, if configured and true
			if config.Skipper != nil && config.Skipper(c) {
				return next(c)
			}

			// Flash a-ah
			shareFlashes(c, config.Inertia)

			return next(c)
		}
	}
}

//
func shareFlashes(c echo.Context, inertia *inertia.Inertia) {
	inertia.Share("flash", map[string]interface{}{})
	if s, err := session.Get("session", c); err == nil {
		inertia.Share("flash", map[string]interface{}{
			"success": s.Flashes("_flash_success"),
			"error":   s.Flashes("_flash_error"),
			"warning": s.Flashes("_flash_warning"),
		})
		_ = s.Save(c.Request(), c.Response())
	}
}
