package handler

import (
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//
func (this *Handler) Organizations(c echo.Context) error {
	if data, err := this.repository.GetOrganizations(); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	} else {
		return this.Render(c, http.StatusOK, "Organizations/Index", map[string]interface{}{
			"organizations": data,
		})
	}
}

//
func (this *Handler) EditOrganization(c echo.Context) error {
	id := this.getAnyParamOrDefault(c, "organization")
	if id != "" {
		id, _ := strconv.Atoi(id)
		if data, err := this.repository.GetOrganizationById(uint(id)); err != nil {
			return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
		} else {
			return this.Render(c, http.StatusOK, "Organizations/Edit", map[string]interface{}{
				"data": data,
			})
		}
	}
	return util.NewError().JSON(c, http.StatusNotFound)
}

//
func (this *Handler) CreateOrganization(c echo.Context) error {
	return this.Render(c, http.StatusOK, "Organizations/Create", map[string]interface{}{})
}

//
func (this *Handler) UpdateOrganization(c echo.Context) error {
	m := model.Organization{}
	if err := this.bindRequest(c, &m); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	// No id, no update
	if m.Id <= 0 {
		return util.NewError().JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.UpdateOrganization(&m)
	if err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	return this.Redirect(c, "/organizations", http.StatusFound, "GET")
}

//
func (this *Handler) StoreOrganization(c echo.Context) error {
	m := model.Organization{}
	if err := this.bindRequest(c, &m); err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	//
	err := this.repository.CreateOrganization(&m)
	if err != nil {
		return util.NewError().AddError(err).JSON(c, http.StatusUnprocessableEntity)
	}
	return this.Redirect(c, "/organizations", http.StatusFound, "GET")
}

//
func (this *Handler) DeleteOrganization(c echo.Context) error {
	m := model.Organization{}
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
	return this.Redirect(c, "/organizations", http.StatusFound, "GET")
}
