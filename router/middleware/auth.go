package middleware

import (
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/service"
	"elipzis.com/inertia-echo/util"
	"fmt"
	"github.com/labstack/echo-contrib/session"
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

			// Check for an authorization header
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				// Fall back to session based JWT token, if non is given
				sess, err := session.Get("session", c)
				if err == nil {
					if user, ok := sess.Values["user"]; ok {
						fmt.Println("TOKEN", *user.(*model.User).Token)
						// Set the JWT Token as header to "fool" the JWT Middleware
						c.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", *user.(*model.User).Token))
					}
				}
			}

			// Fire the Echo JWT first
			jwtFunc := jwtMiddleware(next)
			if err := jwtFunc(c); err != nil {
				// c.Error(err)
				fmt.Println(err)

				switch v := err.(type) {
				case *echo.HTTPError:
					if v.Code == http.StatusUnauthorized {
						// Redirect to login in case something wrong happened while checking the url
						url := util.GetBaseUrl(c)
						c.Request().Method = http.MethodGet
						// Try to find a route named "login"
						for _, route := range c.Echo().Routes() {
							if route.Name == "login" {
								return c.Redirect(http.StatusSeeOther, url+route.Path)
							}
						}
						// Otherwise fall back to a constructed
						return c.Redirect(http.StatusSeeOther, url+"/login")
					}
				}
			}

			// return nil
			return next(c)
		}
	}
}
