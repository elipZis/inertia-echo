package service

import (
	"elipzis.com/inertia-echo/repository/model"
	"errors"
)

//
func (this *Service) Login(email string, password string) (*model.User, error) {
	var user *model.User
	user, err := this.repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user.error.not_found")
	}
	if !user.ValidatePassword(password) {
		return nil, errors.New("user.error.password_mismatch")
	}

	// Generate a token to identify the user
	user.Token = this.GenerateToken(user)
	return user, this.repository.UpdateUser(user)
}

//
func (this *Service) Register(user *model.User) error {
	// Store for later
	email := user.Email
	password := user.Password

	// Create
	if err := this.repository.CreateUser(user); err != nil {
		return err
	}

	// Login the newly created user, implicitly validating that it worked
	user, err := this.Login(email, password)
	return err
}
