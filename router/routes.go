package router

import (
	"elipzis.com/inertia-echo/handler"
	"elipzis.com/inertia-echo/inertia"
	"elipzis.com/inertia-echo/router/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Register the routes
func (router *Router) Register(rootGroup *echo.Group) {
	// Handler
	controller := handler.NewHandler(router.Echo)
	router.Echo.Renderer = controller
	router.Echo.Use(inertia.MiddlewareWithConfig(inertia.MiddlewareConfig{
		Inertia: controller.Inertia,
	}))
	router.Echo.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		_ = controller.Inertia.Render("Error", map[string]interface{}{
			"status": code,
		}).ToResponse(c)
	}

	// Authentication Routes
	rootGroup.GET("/login", controller.LoginForm).Name = "login"
	rootGroup.POST("/login", controller.Login).Name = "login.attempt"
	rootGroup.POST("/register", controller.Register).Name = "register"

	// Index
	rootGroup.GET("/", controller.DashboardIndex, middleware.AuthMiddlewareWithConfig(middleware.AuthMiddlewareConfig{})).Name = "dashboard"
	rootGroup.GET("/dashboard", controller.DashboardIndex)
	rootGroup.GET("/organizations", controller.OrganizationsIndex).Name = "organizations"

	// User handling
	// userGroup := rootGroup.Group("/user")
	// userGroup.Use(jwtMiddleware)
	// userGroup.GET("", controller.GetUser)
	// userGroup.GET("/:id", controller.GetUser)
	// userGroup.PUT("", controller.UpdateUser)

	// Do stuff!
}
