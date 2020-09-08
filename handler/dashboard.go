package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//
func (this *Handler) Dashboard(c echo.Context) error {
	return this.Render(c, http.StatusOK, "Dashboard/Index", map[string]interface{}{})
}
