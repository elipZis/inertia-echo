package repository

import (
	"github.com/jinzhu/gorm"
)

//
type Repository struct {
	Conn *gorm.DB
}

// Create a repository with the Database connection provided
func NewRepository(conn *gorm.DB) (this *Repository) {
	this = new(Repository)
	this.Conn = conn
	return this
}

// Generic storing or updating
func (this *Repository) StoreModel(model interface{}) error {
	if this.Conn.NewRecord(model) {
		return this.CreateModel(model)
	}
	return this.UpdateModel(model)
}

// Create the model and return a resulting error, if any
func (this *Repository) CreateModel(model interface{}) error {
	return this.Conn.Create(model).Error
}

// Update the model and return a resulting error, if any
func (this *Repository) UpdateModel(model interface{}) error {
	return this.Conn.Model(model).Update(model).Error
}

//
func (this *Repository) SaveModel(model interface{}) error {
	return this.Conn.Save(model).Error
}
