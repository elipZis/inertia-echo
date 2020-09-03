package repository

import (
	"elipzis.com/inertia-echo/repository/model"
	"errors"
)

//
func (this *Repository) UpdateUser(model *model.User) error {
	return this.UpdateModel(model)
}

//
func (this *Repository) CreateUser(model *model.User) error {
	// Check that the email is unique
	var count int
	if this.Conn.Model(model).Where("email = ?", model.Email).Count(&count); count > 0 {
		return errors.New("user.error.email_exists")
	}

	return this.CreateModel(model)
}

//
func (this *Repository) GetUserByID(id uint) (*model.User, error) {
	var m model.User
	if err := this.Conn.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

//
func (this *Repository) GetUserByEmail(email string) (*model.User, error) {
	var m model.User
	if err := this.Conn.Where(&model.User{Email: email}).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

//
func (this *Repository) getModelByUserId(id uint, userModel model.UserModel) (*model.UserModel, error) {
	userModel.SetUserId(id)
	if err := this.Conn.Where(&userModel).First(&userModel).Error; err != nil {
		return nil, err
	}
	return &userModel, nil
}
