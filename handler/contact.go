package handler

import (
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//
func (this *Handler) Contacts(c echo.Context) error {
	if data, err := this.repository.GetContacts(); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	} else {
		return this.Render(c, http.StatusOK, "Contacts/Index", map[string]interface{}{
			"contacts": data,
		})
	}
}

//
func (this *Handler) EditContact(c echo.Context) error {
	id := this.getAnyParamOrDefault(c, "contact")
	if id != "" {
		id, _ := strconv.Atoi(id)
		if user, err := this.repository.GetContactById(uint(id)); err != nil {
			return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
		} else {
			return this.Render(c, http.StatusOK, "Contacts/Edit", map[string]interface{}{
				"data": user,
			})
		}
	}
	return util.NewError().JSON(c, http.StatusNotFound)
}

//
func (this *Handler) CreateContact(c echo.Context) error {
	return this.Render(c, http.StatusOK, "Contacts/Create", map[string]interface{}{})
}

//
func (this *Handler) UpdateContact(c echo.Context) error {
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
	return this.Redirect(c, "/organizations", http.StatusFound, "GET")
}

//
func (this *Handler) StoreContact(c echo.Context) error {
	user := model.User{}
	if err := this.bindRequest(c, &user); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.CreateUser(&user)
	if err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	return this.Redirect(c, "/contacts", http.StatusFound, "GET")
}

//
func (this *Handler) DeleteContact(c echo.Context) error {
	user := model.User{}
	if err := this.bindRequest(c, &user); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	// No id, no delete
	if user.Id <= 0 {
		return util.NewError().JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.DeleteModel(&user)
	if err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	return this.Redirect(c, "/contacts", http.StatusFound, "GET")
}
