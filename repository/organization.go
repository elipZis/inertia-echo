package repository

import (
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
func (this *Repository) GetOrganizations() (*[]model.Organization, error) {
	var m []model.Organization
	if err := this.Conn.Find(&m).Error; err != nil {
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
