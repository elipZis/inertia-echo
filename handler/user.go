package handler

import (
	"elipzis.com/inertia-echo/repository/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//
func (this *Handler) Register(c echo.Context) error {
	user := model.User{}
	if err := this.bindAndValidateRequest(c, &user); err != nil {
		return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
	}
	if err := this.service.Register(&user); err != nil {
		return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
	}
	return c.JSON(http.StatusCreated, user)
}

//
func (this *Handler) LoginForm(c echo.Context) error {
	return this.Inertia.Render("Auth/Login", map[string]interface{}{}).ToResponse(c)
}

//
func (this *Handler) Login(c echo.Context) error {
	user := model.User{}
	if err := this.bindAndValidateRequest(c, &user); err != nil {
		// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		// return c.JSON(http.StatusUnprocessableEntity, err)
		return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
	}
	loggedInUser, err := this.service.Login(user.Email, user.Password)
	if err != nil {
		return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
	}
	return c.JSON(http.StatusOK, loggedInUser)
}

//
func (this *Handler) GetUser(c echo.Context) error {
	var id uint
	param, err := strconv.Atoi(c.Param("id"))
	id = uint(param)
	if id == 0 || err != nil {
		userId, err := this.getUserIdFromContext(c)
		if err != nil {
			return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
		}
		id = userId
	}
	// Get user and return
	user, err := this.repository.GetUserByID(id)
	if err != nil {
		return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
	}
	return c.JSON(http.StatusOK, user)
}

//
func (this *Handler) UpdateUser(c echo.Context) error {
	user := model.User{}
	if err := this.bindRequest(c, &user); err != nil {
		return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
	}
	// Store the new user data
	if user.Id <= 0 {
		id, err := this.getUserIdFromContext(c)
		if err != nil {
			return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
		}
		user.Id = id
	}
	//
	err := this.repository.UpdateUser(&user)
	if err != nil {
		return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
	}
	return c.JSON(http.StatusOK, user)
}
