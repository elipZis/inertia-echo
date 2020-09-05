package handler

import (
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/util"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

//
func (this *Handler) LoginForm(c echo.Context) error {
	user, err := this.getUserFromContext(c)
	fmt.Println(user)
	fmt.Println(err)
	return c.Render(200, "Auth/Login", map[string]interface{}{})
	// return this.Inertia.Render("Auth/Login", map[string]interface{}{}).ToResponse(c)
}

//
func (this *Handler) Login(c echo.Context) error {
	user := model.User{}
	if err := this.bindAndValidateRequest(c, &user); err != nil {
		// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		// return c.JSON(http.StatusUnprocessableEntity, err)
		// return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
		return util.NewError().AddError(err).Render(c, "Auth/Login", http.StatusUnprocessableEntity)
	}
	loggedInUser, err := this.service.Login(user.Email, user.Password)
	if err != nil {
		// return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	return c.JSON(http.StatusOK, loggedInUser)
}

//
func (this *Handler) Logout(c echo.Context) error {
	return c.Render(200, "Auth/Login", map[string]interface{}{})
	// return this.Inertia.Render("Auth/Login", map[string]interface{}{}).ToResponse(c)
}
