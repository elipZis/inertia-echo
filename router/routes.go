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
		c.Request().Method = http.MethodGet
		return c.Redirect(http.StatusMovedPermanently, util.GetRedirectUrl(c, "/dashboard"))
	})
	rootGroup.GET("/login", controller.LoginForm).Name = "login"
	rootGroup.POST("/login", controller.Login).Name = "login.attempt"
	rootGroup.POST("/logout", controller.Logout).Name = "logout"

	// Dashboard
	dashboardGroup := rootGroup.Group("/dashboard")
	dashboardGroup.Use(middleware.AuthMiddlewareWithConfig(middleware.AuthMiddlewareConfig{}))
	dashboardGroup.GET("", controller.Dashboard).Name = "dashboard"

	// Organizations
	organizationsGroup := rootGroup.Group("/organizations")
	organizationsGroup.Use(middleware.AuthMiddlewareWithConfig(middleware.AuthMiddlewareConfig{}))
	organizationsGroup.GET("", controller.Organizations).Name = "organizations"
	organizationsGroup.GET("/create", controller.CreateOrganization).Name = "organizations.create"
	organizationsGroup.POST("/store", controller.StoreOrganization).Name = "organizations.store"
	organizationsGroup.GET("/edit", controller.EditOrganization).Name = "organizations.edit"
	organizationsGroup.POST("/update", controller.UpdateOrganization).Name = "organizations.update"
	organizationsGroup.DELETE("/delete", controller.DeleteOrganization).Name = "organizations.destroy"

	// User handling
	usersGroup := rootGroup.Group("/users")
	usersGroup.Use(middleware.AuthMiddlewareWithConfig(middleware.AuthMiddlewareConfig{}))
	usersGroup.GET("", controller.Users).Name = "users"
	usersGroup.GET("/edit", controller.EditUser).Name = "users.edit"
	usersGroup.GET("/create", controller.CreateUser).Name = "users.create"
	usersGroup.POST("/update", controller.UpdateUser).Name = "users.update"
	usersGroup.POST("/store", controller.StoreUser).Name = "users.store"
	usersGroup.DELETE("/delete", controller.DeleteUser).Name = "users.delete"

	// Contacts
	contactsGroup := rootGroup.Group("/contacts")
	contactsGroup.Use(middleware.AuthMiddlewareWithConfig(middleware.AuthMiddlewareConfig{}))
	contactsGroup.GET("", controller.Contacts).Name = "contacts"
	contactsGroup.GET("/create", controller.CreateContact).Name = "contacts.create"
	contactsGroup.POST("/store", controller.StoreContact).Name = "contacts.store"
	contactsGroup.GET("/edit", controller.EditContact).Name = "contacts.edit"
	contactsGroup.POST("/update", controller.UpdateContact).Name = "contacts.update"
	contactsGroup.DELETE("/delete", controller.DeleteContact).Name = "contacts.destroy"

	// Reports
	reportsGroup := rootGroup.Group("/reports")
	reportsGroup.Use(middleware.AuthMiddlewareWithConfig(middleware.AuthMiddlewareConfig{}))
	reportsGroup.GET("", controller.Reports).Name = "reports"

	// 500 error
	rootGroup.GET("/500", func(c echo.Context) error {
		return c.HTML(http.StatusInternalServerError, "500er Error")
	}).Name = "500"

	// Do stuff!
}
