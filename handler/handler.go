package handler

import (
	"elipzis.com/inertia-echo/repository"
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/service"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

//
type Handler struct {
	echo       *echo.Echo
	service    *service.Service
	repository *repository.Repository

	// templates *template.Template
	// Inertia   *inertia.Inertia
}

//
func NewHandler(echo *echo.Echo) (this *Handler) {
	this = new(Handler)
	this.echo = echo
	this.repository = repository.NewRepository(repository.DB.Conn)
	this.service = service.NewService(this.repository)

	// this.Inertia = inertia.NewInertia(echo)
	// this.Inertia.SetMixVersion()
	return this
}

//
func (this *Handler) getAnyParamOrDefault(c echo.Context, field string, defaultValue ...interface{}) interface{} {
	value := c.FormValue(field)
	if value == "" {
		value = c.QueryParam(field)
		if value == "" {
			value = c.Param(field)
		}
	}

	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return value
}

//
func (this *Handler) getUserFromContext(c echo.Context) (*model.User, error) {
	tokenUser, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("jwt.error.no_user")
	}
	claims := tokenUser.Claims.(*service.JWTCustomClaims)
	return &model.User{
		Id:    claims.Id,
		Email: claims.Email,
		Token: &tokenUser.Raw,
	}, nil
}

//
func (this *Handler) getUserIdFromContext(c echo.Context) (uint, error) {
	user, err := this.getUserFromContext(c)
	return user.Id, err
}

// Generic binding of requests to models
func (this *Handler) bindRequest(c echo.Context, model interface{}) error {
	if err := c.Bind(model); err != nil {
		return err
	}
	return nil
}

// Generic validation of requests against go.validate
func (this *Handler) validateRequest(c echo.Context, dto interface{}) error {
	if err := c.Validate(dto); err != nil {
		return err
	}
	switch dto.(type) {
	case *model.UserModel:
		return this.validateAndBindUser(c, dto.(model.UserModel))
	}
	return nil
}

// Validate that a given model userid matches the authenticated user
func (this *Handler) validateAndBindUser(c echo.Context, model model.UserModel) error {
	id, err := this.getUserIdFromContext(c)
	if err != nil || (model.GetUserId() != nil && model.GetUserId() != &id) {
		return errors.New("handler.error.user_mismatch")
	}
	model.SetUserId(id)
	return nil
}

// Generic binding of requests to models and validating against go.validate
func (this *Handler) bindAndValidateRequest(c echo.Context, model interface{}) error {
	if err := this.bindRequest(c, model); err != nil {
		return err
	}
	if err := this.validateRequest(c, model); err != nil {
		return err
	}
	return nil
}
