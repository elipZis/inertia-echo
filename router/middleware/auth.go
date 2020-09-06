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
	"reflect"
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

			sess, err := session.Get("session", c)
			if err == nil {
				if user, ok := sess.Values["user"]; ok {
					fmt.Println("TOKEN", *user.(*model.User).Token)
					// Set the JWT Token as header to "fool" the JWT Middleware
					c.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", *user.(*model.User).Token))
				}
			}

			fmt.Println(c.Request().Header)

			// Fire the Echo JWT first
			jwtFunc := jwtMiddleware(next)
			if err := jwtFunc(c); err != nil {
				fmt.Println("Auth error", err)
				fmt.Println(reflect.TypeOf(err))
				fmt.Println(err.(*echo.HTTPError).Internal)
				fmt.Println(err.(*echo.HTTPError).Code)
				fmt.Println(err.(*echo.HTTPError).Message)
				fmt.Println(err.(*echo.HTTPError).Unwrap())
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

			return nil
			return next(c)
		}
	}
}
