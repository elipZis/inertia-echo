package handler

import (
	"github.com/labstack/echo/v4"
)

//
func (this *Handler) DashboardIndex(c echo.Context) error {
	// sess, _ := session.Get("session", c)
	// sess.Options = &sessions.Options{
	// 	Path:     "/",
	// 	MaxAge:   86400 * 7,
	// 	HttpOnly: true,
	// }
	// sess.Values["auth"] = "user"
	// sess.Save(c.Request(), c.Response())
	// return this.Inertia.Render("Dashboard/Index", map[string]interface{}{}).ToResponse(c)
	return c.Render(200, "Dashboard/Index", map[string]interface{}{})
	// return c.JSON(http.StatusOK, []string{"Hello"})
}

func (this *Handler) OrganizationsIndex(c echo.Context) error {
	return c.Render(200, "Dashboard/Organizations", nil)
	// return this.Inertia.Render("Dashboard/Organizations", nil).ToResponse(c)
	// return c.JSON(http.StatusOK, []string{"Hello"})
}
