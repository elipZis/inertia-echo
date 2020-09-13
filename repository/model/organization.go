package model

import (
	"elipzis.com/inertia-echo/repository/model/trait"
)

//
type Organization struct {
	Id   uint   `gorm:"primary_key"`
	Name string `gorm:"not null" validate:"required"`

	Email      string `gorm:"null" validate:"email"`
	Phone      string `gorm:"null" validate:"required"`
	Address    string `gorm:"null"`
	City       string `gorm:"null"`
	Region     string `gorm:"null"`
	Country    string `gorm:"null" validate:"omitempty,alpha,len=2"`
	PostalCode string `gorm:"null"`

	Contacts []Contact `json:",omitempty" gorm:"foreignKey:OrganizationId"`

	trait.Timestampable
}
