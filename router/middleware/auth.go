package middleware

import (
	"elipzis.com/inertia-echo/service"
	"elipzis.com/inertia-echo/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

//
type AuthMiddlewareConfig struct {
	Skipper   middleware.Skipper
	JwtConfig *middleware.JWTConfig
}

//
var DefaultJWTConfig = middleware.JWTConfig{
	Claims:     &service.JWTCustomClaims{},
	SigningKey: service.JWTSecret,
}

//
func AuthMiddlewareWithConfig(config AuthMiddlewareConfig) echo.MiddlewareFunc {
	if config.JwtConfig == nil {
		config.JwtConfig = &DefaultJWTConfig
	}
	//
	var jwtMiddleware = middleware.JWTWithConfig(*config.JwtConfig)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip, if configured and true
			if config.Skipper != nil && config.Skipper(c) {
				return next(c)
			}

			// Fire the Echo JWT first
			jwtFunc := jwtMiddleware(next)
			if err := jwtFunc(c); err != nil {
				// c.Error(err)

				// Redirect to login in case something wrong happened while checking the url
				url := util.GetBaseUrl(c)
				// Try to find a route named "login"
				for _, route := range c.Echo().Routes() {
					if route.Name == "login" {
						return c.Redirect(http.StatusTemporaryRedirect, url+route.Path)
					}
				}
				// Otherwise fall back to a constructed
				return c.Redirect(http.StatusTemporaryRedirect, url+"/login")
			}

			return next(c)
		}
	}
}
