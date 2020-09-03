package router

import (
	"context"
	"elipzis.com/inertia-echo/service"
	"fmt"
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"strconv"
)

//
type Router struct {
	Echo *echo.Echo
}

// Create a new Echo router and configure some middlewares
func NewRouter() (this *Router) {
	this = new(Router)
	this.Echo = echo.New()

	// Logging
	this.Echo.Logger.SetLevel(log.WARN)
	if debug, _ := strconv.ParseBool(service.GetEnvOrDefault("DEBUG", "false")); debug {
		this.Echo.Logger.SetLevel(log.DEBUG)
	}

	// Global middlewares
	this.Echo.Pre(middleware.RemoveTrailingSlash())
	this.Echo.Use(middleware.Logger())
	this.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	this.Echo.Use(session.Middleware(sessions.NewCookieStore([]byte(service.GetEnvOrDefault("SESSION_SECRET", "supersecretsessionsecret")))))
	this.Echo.Static("/", "public")

	// go-playground/validation
	this.Echo.Validator = NewValidator()
	return this
}

// Start the server
func (this *Router) Run() {
	// Start server
	go func() {
		if err := this.Echo.Start(fmt.Sprintf("%s:%s", service.GetEnvOrDefault("HOST", "localhost"), service.GetEnvOrDefault("PORT", "1323"))); err != nil {
			this.Echo.Logger.Info("Shutting down the server...")
		}
	}()
}

//
func (this *Router) Shutdown(ctx context.Context) {
	// return this.Echo.Shutdown(ctx)
	if err := this.Echo.Shutdown(ctx); err != nil {
		this.Echo.Logger.Fatal(err)
	}
}
