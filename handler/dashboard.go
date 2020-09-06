package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

//
func (this *Handler) DashboardIndex(c echo.Context) error {
	fmt.Println("Dash")
	return this.Render(c, 200, "Dashboard/Index", map[string]interface{}{})
	// return c.JSON(http.StatusOK, []string{"Hello"})
}

func (this *Handler) OrganizationsIndex(c echo.Context) error {
	fmt.Println("Org")
	return this.Render(c, 200, "Dashboard/Organizations", map[string]interface{}{})
	// return this.Inertia.Render("Dashboard/Organizations", nil).ToResponse(c)
	// return c.JSON(http.StatusOK, []string{"Hello"})
}
