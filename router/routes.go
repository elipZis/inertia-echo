package router

import (
	"elipzis.com/inertia-echo/handler"
	"elipzis.com/inertia-echo/inertia"
	"elipzis.com/inertia-echo/router/middleware"
	"elipzis.com/inertia-echo/util"
	"github.com/labstack/echo/v4"
)

// Register the routes
func (this *Router) Register(rootGroup *echo.Group) {
	// Handler
	controller := handler.NewHandler(this.Echo)
	// router.Echo.Renderer = controller
	this.Echo.Use(inertia.MiddlewareWithConfig(inertia.MiddlewareConfig{
		Inertia: controller.Inertia,
	}))

	// Authentication Routes
	rootGroup.GET("", func(c echo.Context) error {
		return c.Redirect(302, util.GetRedirectUrl(c, "/dashboard"))
	})
	rootGroup.GET("/login", controller.LoginForm).Name = "login"
	rootGroup.POST("/login", controller.Login).Name = "login.attempt"
	rootGroup.POST("/logout", controller.Logout).Name = "logout"

	// Dashboard
	dashboardGroup := rootGroup.Group("/dashboard")
	dashboardGroup.Use(middleware.AuthMiddlewareWithConfig(middleware.AuthMiddlewareConfig{}))
	dashboardGroup.GET("", controller.DashboardIndex).Name = "dashboard"

	// Organizations
	organizationsGroup := rootGroup.Group("/organizations")
	organizationsGroup.Use(middleware.AuthMiddlewareWithConfig(middleware.AuthMiddlewareConfig{}))
	organizationsGroup.GET("", controller.OrganizationsIndex).Name = "organizations"

	// User handling
	usersGroup := rootGroup.Group("/users")
	usersGroup.Use(middleware.AuthMiddlewareWithConfig(middleware.AuthMiddlewareConfig{}))
	usersGroup.GET("", controller.GetUser).Name = "users"
	usersGroup.GET("/edit", controller.UpdateUser).Name = "users.edit"
	// usersGroup.PUT("", controller.UpdateUser).Name = "dashboard"

	// Do stuff!
}
