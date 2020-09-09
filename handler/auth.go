package handler

import (
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

//
func (this *Handler) Test(c echo.Context) error {
	return c.Render(http.StatusOK, c.QueryParam("template"), map[string]interface{}{
		"test": "ok",
	})
}

//
func (this *Handler) LoginForm(c echo.Context) error {
	return c.Render(http.StatusOK, "Auth/Login", map[string]interface{}{})
}

//
func (this *Handler) Login(c echo.Context) error {
	user := model.User{}
	if err := this.bindAndValidateRequest(c, &user); err != nil {
		return util.NewError().AddError(err).Render(c, "Auth/Login", http.StatusUnprocessableEntity)
	}
	loggedInUser, err := this.service.Login(user.Email, user.Password)
	if err != nil {
		return util.NewError().AddError(err).Render(c, "Auth/Login", http.StatusUnprocessableEntity)
	}
	if err = this.setSession(c, "user", loggedInUser, nil); err != nil {
		return util.NewError().AddError(err).Render(c, "Auth/Login", http.StatusUnprocessableEntity)
	}
	return c.Redirect(http.StatusFound, "/")
}

//
func (this *Handler) Logout(c echo.Context) error {
	if err := this.deleteSession(c); err != nil {
		return util.NewError().AddError(err).Render(c, "Auth/Login", http.StatusUnprocessableEntity)
	}
	return c.Render(200, "Auth/Login", map[string]interface{}{})
	// return this.Inertia.Render("Auth/Login", map[string]interface{}{}).ToResponse(c)
}

//
func (this *Handler) Register(c echo.Context) error {
	user := model.User{}
	if err := this.bindAndValidateRequest(c, &user); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	if err := this.service.Register(&user); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	return c.JSON(http.StatusCreated, user)
}
