package model

import (
	"elipzis.com/inertia-echo/repository/model/trait"
)

//
type Contact struct {
	Id        uint   `gorm:"primary_key"`
	FirstName string `gorm:"not null" validate:"required"`
	LastName  string `gorm:"not null" validate:"required"`

	Email      string `gorm:"null" validate:"email"`
	Phone      string `gorm:"null"`
	Address    string `gorm:"null"`
	City       string `gorm:"null"`
	Region     string `gorm:"null"`
	Country    string `gorm:"null" validate:"alpha;len=2"`
	PostalCode string `gorm:"null"`

	OrganizationId int
	Organization   Organization `gorm:"foreignKey:OrganizationId;references:Id"`

	trait.Timestampable
}