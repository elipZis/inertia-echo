package repository

import (
	dto "elipzis.com/inertia-echo/handler/model"
	"elipzis.com/inertia-echo/repository/model"
)

//
func (this *Repository) UpdateContact(model *model.Contact) error {
	return this.UpdateModel(model)
}

//
func (this *Repository) CreateContact(model *model.Contact) error {
	return this.CreateModel(model)
}

//
func (this *Repository) GetContacts(filter *dto.Filter) (*[]model.Contact, error) {
	if filter == nil {
		filter = &dto.Filter{}
	}
	var m []model.Contact
	if err := this.Conn.Preload("Organization").
		Where("LOWER(first_name) LIKE ? or LOWER(last_name) LIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%").
		Find(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

//
func (this *Repository) GetContactById(id uint) (*model.Contact, error) {
	var m model.Contact
	if err := this.Conn.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
