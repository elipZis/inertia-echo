package repository

import (
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
func (this *Repository) GetContacts() (*[]model.Contact, error) {
	var m []model.Contact
	if err := this.Conn.Find(&m).Error; err != nil {
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
