package router

import (
	"elipzis.com/inertia-echo/handler"
	"github.com/labstack/echo/v4"
)

// Register the routes
func (router *Router) Register(rootGroup *echo.Group) {
	// Handler
	controller := handler.NewHandler(router.Echo)
	// router.Echo.Renderer = controller
	// router.Echo.Use(inertia.Middleware(inertia.MiddlewareConfig{
	// 	Inertia: controller.Inertia,
	// }))
	// router.Echo.HTTPErrorHandler = func(err error, c echo.Context) {
	// 	code := http.StatusInternalServerError
	// 	if he, ok := err.(*echo.HTTPError); ok {
	// 		code = he.Code
	// 	}
	// 	// _ = controller.Inertia.Render("Error", map[string]interface{}{
	// 	// 	"status": code,
	// 	// }).ToResponse(c)
	// 	_ = c.Render(code, "Error", map[string]interface{}{
	// 		"status": code,
	// 	})
	// }

	// Authentication Routes
	rootGroup.GET("/login", controller.LoginForm).Name = "login"
	rootGroup.POST("/login", controller.Login).Name = "login.attempt"
	rootGroup.POST("/register", controller.Register).Name = "register"

	// Index
	// rootGroup.GET("/", controller.DashboardIndex, middleware.AuthMiddlewareWithConfig(middleware.AuthMiddlewareConfig{})).Name = "dashboard"
	rootGroup.GET("/", controller.DashboardIndex).Name = "dashboard"
	rootGroup.GET("/organizations", controller.OrganizationsIndex).Name = "organizations"

	// User handling
	// userGroup := rootGroup.Group("/user")
	// userGroup.Use(jwtMiddleware)
	// userGroup.GET("", controller.GetUser)
	// userGroup.GET("/:id", controller.GetUser)
	// userGroup.PUT("", controller.UpdateUser)

	// Do stuff!
}
