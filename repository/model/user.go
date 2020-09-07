package model

import (
	"elipzis.com/inertia-echo/repository/model/trait"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//
type UserModel interface {
	GetUserId() *uint
	SetUserId(userId uint)
}

//
type User struct {
	Id       uint    `gorm:"primary_key"`
	Email    string  `gorm:"unique_index;not null" validate:"required,email"`
	Password string  `gorm:"not null" validate:"required,min=6"`
	Token    *string `gorm:"null"`

	FirstName string `gorm:"null"`
	LastName  string `gorm:"null"`

	trait.Timestampable
	trait.Softdeleteable
}

//
func (this *User) Name() string {
	return this.FirstName + " " + this.LastName
}

//
func (this *User) BeforeCreate(scope *gorm.Scope) error {
	hashedPassword, _ := this.HashPassword(this.Password)
	scope.SetColumn("Password", hashedPassword)
	return nil
}

//
func (this *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("user.error.password_empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

//
func (this *User) ValidatePassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(this.Password), []byte(plain))
	return err == nil
}
