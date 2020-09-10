package handler

import (
	dto "elipzis.com/inertia-echo/handler/model"
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//
func (this *Handler) Contacts(c echo.Context) error {
	filter := dto.Filter{}
	if err := this.bindRequest(c, &filter); err != nil {
		return this.ErrorResponse(c, err)
	}
	if data, err := this.repository.GetContacts(&filter); err != nil {
		return this.ErrorResponse(c, err)
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
			return this.ErrorResponse(c, err)
		} else {
			organizations, _ := this.repository.GetOrganizations(nil)
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
	organizations, _ := this.repository.GetOrganizations(nil)
	return this.Render(c, http.StatusOK, "Contacts/Create", map[string]interface{}{
		"organizations": organizations,
	})
}

//
func (this *Handler) UpdateContact(c echo.Context) error {
	m := model.Contact{}
	if err := this.bindAndValidateRequest(c, &m); err != nil {
		return this.ErrorResponse(c, err)
	}
	// No id, no update
	if m.Id <= 0 {
		return util.NewError().JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.UpdateContact(&m)
	if err != nil {
		return this.ErrorResponse(c, err)
	}
	return this.addSuccessFlash(c, "Contact updated").Redirect(c, "/contacts", http.StatusFound, "GET")
}

//
func (this *Handler) StoreContact(c echo.Context) error {
	m := model.Contact{}
	if err := this.bindAndValidateRequest(c, &m); err != nil {
		return this.ErrorResponse(c, err)
	}
	//
	err := this.repository.CreateContact(&m)
	if err != nil {
		return this.ErrorResponse(c, err)
	}
	return this.addSuccessFlash(c, "Contact stored").Redirect(c, "/contacts", http.StatusFound, "GET")
}

//
func (this *Handler) DeleteContact(c echo.Context) error {
	m := model.Contact{}
	if err := this.bindRequest(c, &m); err != nil {
		return this.ErrorResponse(c, err)
	}
	// No id, no delete
	if m.Id <= 0 {
		return util.NewError().JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.DeleteModel(&m)
	if err != nil {
		return this.ErrorResponse(c, err)
	}
	return this.addSuccessFlash(c, "Contact deleted").Redirect(c, "/contacts", http.StatusFound, "GET")
}
