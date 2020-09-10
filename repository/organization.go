package repository

import (
	dto "elipzis.com/inertia-echo/handler/model"
	"elipzis.com/inertia-echo/repository/model"
)

//
func (this *Repository) UpdateOrganization(model *model.Organization) error {
	return this.UpdateModel(model)
}

//
func (this *Repository) CreateOrganization(model *model.Organization) error {
	return this.CreateModel(model)
}

//
func (this *Repository) GetOrganizations(filter *dto.Filter) (*[]model.Organization, error) {
	if filter == nil {
		filter = &dto.Filter{}
	}
	var m []model.Organization
	if err := this.Conn.
		Where("LOWER(name) LIKE ?", "%"+filter.Search+"%").
		Find(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

//
func (this *Repository) GetOrganizationById(id uint) (*model.Organization, error) {
	var m model.Organization
	if err := this.Conn.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

//
func (this *Repository) GetOrganizationContacts(organization *model.Organization) *[]model.Contact {
	var contacts []model.Contact
	this.Conn.Model(&organization).Related(&contacts)
	return &contacts
}
