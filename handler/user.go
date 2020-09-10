package handler

import (
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

//
func (this *Handler) Users(c echo.Context) error {
	if users, err := this.repository.GetUsers(); err != nil {
		return this.ErrorResponse(c, err)
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
			return this.ErrorResponse(c, err)
		}
		id = userId
	}
	// Get user and return
	user, err := this.repository.GetUserById(id)
	if err != nil {
		return this.ErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, user)
}

//
func (this *Handler) EditUser(c echo.Context) error {
	id := this.getAnyParamOrDefault(c, "user")
	if id != "" {
		id, _ := strconv.Atoi(id)
		if user, err := this.repository.GetUserById(uint(id)); err != nil {
			return this.ErrorResponse(c, err)
		} else {
			return this.Render(c, http.StatusOK, "Users/Edit", map[string]interface{}{
				"data": user,
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
	if err := this.bindAndValidateRequest(c, &user); err != nil {
		return this.ErrorResponse(c, err)
	}
	// No id, no update
	if user.Id <= 0 {
		return util.NewError().JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.UpdateUser(&user)
	if err != nil {
		return this.ErrorResponse(c, err)
	}
	return this.addSuccessFlash(c, "User updated").Redirect(c, "/users", http.StatusFound, "GET")
}

//
func (this *Handler) StoreUser(c echo.Context) error {
	user := model.User{}
	if err := this.bindAndValidateRequest(c, &user); err != nil {
		return this.ErrorResponse(c, err)
	}
	//
	err := this.repository.CreateUser(&user)
	if err != nil {
		return this.ErrorResponse(c, err)
	}
	return this.addSuccessFlash(c, "User stored").Redirect(c, "/users", http.StatusFound, "GET")
}

//
func (this *Handler) DeleteUser(c echo.Context) error {
	user := model.User{}
	if err := this.bindRequest(c, &user); err != nil {
		return this.ErrorResponse(c, err)
	}
	// No id, no delete
	if user.Id <= 0 {
		return util.NewError().JSON(c, http.StatusUnprocessableEntity)
	}
	//
	t := time.Now()
	user.DeletedAt = &t
	err := this.repository.UpdateUser(&user)
	if err != nil {
		return this.ErrorResponse(c, err)
	}
	return this.addSuccessFlash(c, "User deleted").Redirect(c, "/users", http.StatusFound, "GET")
}
