package router

import (
	"elipzis.com/inertia-echo/handler"
	"elipzis.com/inertia-echo/inertia"
	"elipzis.com/inertia-echo/router/middleware"
	"elipzis.com/inertia-echo/util"
	"github.com/labstack/echo/v4"
	"net/http"
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
		return c.Redirect(http.StatusPermanentRedirect, util.GetRedirectUrl(c, "/dashboard"))
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
	usersGroup.GET("", controller.Users).Name = "users"
	usersGroup.GET("/edit", controller.EditUser).Name = "users.edit"
	usersGroup.GET("/create", controller.CreateUser).Name = "users.create"
	usersGroup.POST("/update", controller.UpdateUser).Name = "users.update"
	usersGroup.POST("/store", controller.StoreUser).Name = "users.store"

	// Do stuff!
}
