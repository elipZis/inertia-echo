package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//
func (this *Handler) Reports(c echo.Context) error {
	return this.Render(c, http.StatusOK, "Reports/Index", map[string]interface{}{})
}
