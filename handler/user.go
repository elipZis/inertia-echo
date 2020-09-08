package handler

import (
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/util"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//
func (this *Handler) Users(c echo.Context) error {
	fmt.Println("FICKT EUCH")
	if users, err := this.repository.GetUsers(); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	} else {
		return this.Render(c, http.StatusOK, "Users/Index", map[string]interface{}{
			"users": users,
		})
	}
}

//
func (this *Handler) GetUser(c echo.Context) error {
	var id uint
	param, err := strconv.Atoi(c.Param("id"))
	id = uint(param)
	if id == 0 || err != nil {
		userId, err := this.getUserIdFromContext(c)
		if err != nil {
			return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
			// return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
		}
		id = userId
	}
	// Get user and return
	user, err := this.repository.GetUserByID(id)
	if err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
		// return this.createErrorResponse(c, err, http.StatusUnprocessableEntity)
	}
	return c.JSON(http.StatusOK, user)
}

//
func (this *Handler) EditUser(c echo.Context) error {
	userId := this.getAnyParamOrDefault(c, "user")
	if userId != "" {
		id, _ := strconv.Atoi(userId)
		if user, err := this.repository.GetUserByID(uint(id)); err != nil {
			return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
		} else {
			return this.Render(c, http.StatusOK, "Users/Edit", map[string]interface{}{
				"user": user,
			})
		}
	}
	return util.NewError().JSON(c, http.StatusNotFound)
}

//
func (this *Handler) CreateUser(c echo.Context) error {
	return this.Render(c, http.StatusOK, "Users/Create", map[string]interface{}{})
}

//
func (this *Handler) UpdateUser(c echo.Context) error {
	user := model.User{}
	if err := this.bindRequest(c, &user); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	// No id, no update
	if user.Id <= 0 {
		return util.NewError().JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.UpdateUser(&user)
	if err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	return c.Redirect(http.StatusOK, "/users")
}

//
func (this *Handler) StoreUser(c echo.Context) error {
	user := model.User{}
	if err := this.bindRequest(c, &user); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.CreateUser(&user)
	if err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	fmt.Println("STORE REDIRECT")
	fmt.Println(util.GetRedirectUrl(c, "/users"))
	c.Request().Method = http.MethodGet
	return c.Redirect(http.StatusFound, util.GetRedirectUrl(c, "/users"))
}
