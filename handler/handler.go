package handler

import (
	"elipzis.com/inertia-echo/inertia"
	"elipzis.com/inertia-echo/repository"
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/service"
	"encoding/gob"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

//
type Handler struct {
	echo       *echo.Echo
	service    *service.Service
	repository *repository.Repository
	Inertia    *inertia.Inertia
}

//
func NewHandler(echo *echo.Echo) (this *Handler) {
	this = new(Handler)
	this.echo = echo
	this.repository = repository.NewRepository(repository.DB.Conn)
	this.service = service.NewService(this.repository)
	this.Inertia = inertia.NewInertia(echo)
	return this
}

//
func (this *Handler) Render(c echo.Context, code int, name string, data map[string]interface{}) error {
	if user, err := this.getUserFromContext(c); err == nil {
		data["user"] = user
	}
	return c.Render(code, name, data)
}

//
func (this *Handler) Redirect(c echo.Context, path string, code int, method string) error {
	if code == 0 {
		code = http.StatusFound
	}
	// Change of request method does not work with 307 and 308 by specification (as it seems)
	if code >= http.StatusTemporaryRedirect && code <= http.StatusPermanentRedirect && method != "" {
		code = http.StatusFound
	}

	if method == "" {
		method = http.MethodGet
		c.Request().Method = method
	} else {
		c.Request().Method = method
	}
	return c.Redirect(code, GetRedirectUrl(c, path))
}

//
func (this *Handler) getAnyParamOrDefault(c echo.Context, field string, defaultValue ...string) string {
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
		Id:        claims.Id,
		Email:     claims.Email,
		FirstName: claims.FirstName,
		LastName:  claims.LastName,
		Token:     tokenUser.Raw,
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

//
func (this *Handler) setSession(c echo.Context, key string, value interface{}, options *sessions.Options) error {
	// Register whatever object to be saved
	gob.Register(value)

	// Set it
	s, _ := session.Get("session", c)
	if options != nil {
		s.Options = options
	}
	s.Values[key] = value
	return s.Save(c.Request(), c.Response())
}

//
func (this *Handler) deleteSession(c echo.Context) error {
	s, _ := session.Get("session", c)
	s.Options.MaxAge = -1
	return s.Save(c.Request(), c.Response())
}

//
func (this *Handler) addFlash(c echo.Context, value interface{}, vars ...string) error {
	s, _ := session.Get("session", c)
	s.AddFlash(value, vars...)
	return s.Save(c.Request(), c.Response())
}

//
func (this *Handler) addSuccessFlash(c echo.Context, value interface{}) error {
	return this.addFlash(c, value, "_flash_success")
}

//
func (this *Handler) addErrorFlash(c echo.Context, value interface{}) error {
	return this.addFlash(c, value, "_flash_error")
}

//
func (this *Handler) addWarningFlash(c echo.Context, value interface{}) error {
	return this.addFlash(c, value, "_flash_warning")
}

//
func GetRedirectUrl(c echo.Context, path string) string {
	return GetBaseUrl(c) + path
}

//
func GetBaseUrl(c echo.Context) string {
	req, scheme := c.Request(), c.Scheme()
	host := req.Host
	url := scheme + "://" + host // + req.RequestURI
	return strings.TrimSuffix(url, "/")
}
