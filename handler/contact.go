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
		if data, err := this.repository.GetContactById(uint(id)); err != nil {
			return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
		} else {
			organizations, _ := this.repository.GetOrganizations()
			return this.Render(c, http.StatusOK, "Contacts/Edit", map[string]interface{}{
				"data":          data,
				"organizations": organizations,
			})
		}
	}
	return util.NewError().JSON(c, http.StatusNotFound)
}

//
func (this *Handler) CreateContact(c echo.Context) error {
	organizations, _ := this.repository.GetOrganizations()
	return this.Render(c, http.StatusOK, "Contacts/Create", map[string]interface{}{
		"organizations": organizations,
	})
}

//
func (this *Handler) UpdateContact(c echo.Context) error {
	m := model.Contact{}
	if err := this.bindAndValidateRequest(c, &m); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	// No id, no update
	if m.Id <= 0 {
		return util.NewError().JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.UpdateContact(&m)
	if err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	return this.Redirect(c, "/organizations", http.StatusFound, "GET")
}

//
func (this *Handler) StoreContact(c echo.Context) error {
	m := model.Contact{}
	if err := this.bindRequest(c, &m); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.CreateContact(&m)
	if err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	return this.Redirect(c, "/contacts", http.StatusFound, "GET")
}

//
func (this *Handler) DeleteContact(c echo.Context) error {
	m := model.Contact{}
	if err := this.bindRequest(c, &m); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	// No id, no delete
	if m.Id <= 0 {
		return util.NewError().JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.DeleteModel(&m)
	if err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	return this.Redirect(c, "/contacts", http.StatusFound, "GET")
}
