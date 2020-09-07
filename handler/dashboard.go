package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//
func (this *Handler) DashboardIndex(c echo.Context) error {
	return this.Render(c, http.StatusOK, "Dashboard/Index", map[string]interface{}{})
}

//
func (this *Handler) OrganizationsIndex(c echo.Context) error {
	return this.Render(c, http.StatusOK, "Dashboard/Organizations", map[string]interface{}{})
}
