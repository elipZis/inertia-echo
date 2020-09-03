package handler

import (
	"elipzis.com/inertia-echo/inertia"
	"elipzis.com/inertia-echo/repository"
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/service"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//
type Handler struct {
	echo       *echo.Echo
	service    *service.Service
	repository *repository.Repository

	templates *template.Template
	Inertia   *inertia.Inertia
}

//
func NewHandler(echo *echo.Echo) (this *Handler) {
	this = new(Handler)
	this.echo = echo
	this.repository = repository.NewRepository(repository.DB.Conn)
	this.service = service.NewService(this.repository)

	this.Inertia = inertia.NewInertia()
	this.Inertia.SetMixVersion()
	this.templates = template.Must(template.New("").Funcs(template.FuncMap{
		"inertia": func(v interface{}) template.HTML {
			retVal, _ := json.Marshal(v)
			return template.HTML(fmt.Sprintf("<div id='app' data-page='%s'></div>", string(retVal)))
		},
		"json_encode": func(v interface{}) template.JS {
			retVal, _ := json.Marshal(v)
			return template.JS(string(retVal))
		},
		"json_encode_raw": func(v interface{}) string {
			retVal, _ := json.Marshal(v)
			return string(retVal)
		},
		"routes": func() template.JS {
			this.echo.Routes()
			retVal, _ := json.Marshal(this.echo.Routes())
			return template.JS(string(retVal))
		},
		"routes_ziggy": func(v interface{}) template.HTML {
			ziggy := inertia.NewZiggy(this.echo, v.(map[string]interface{}))
			retVal, _ := json.Marshal(ziggy)
			return template.HTML(fmt.Sprintf("<script>var Ziggy = %s;</script>", string(retVal)))
		},
		"mix": func(path string) template.HTML {
			if !strings.HasPrefix(path, "/") {
				path = "/" + path
			}

			manifestFile, err := os.Open(service.GetEnvOrDefault("PUBLIC_PATH", "public") + "/mix-manifest.json")
			defer manifestFile.Close()
			if err != nil {
				return template.HTML(path)
			}

			manifestData, err := ioutil.ReadAll(manifestFile)
			if err != nil {
				return template.HTML(path)
			}
			var manifest map[string]string
			if err = json.Unmarshal(manifestData, &manifest); err != nil {
				return template.HTML(path)
			}

			return template.HTML(manifest[path])
		},
	}).ParseGlob(service.GetEnvOrDefault("RESOURCES_PATH", "resources") + "/views/*.html"))
	return this
}

// Render renders a template document
func (this *Handler) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return this.templates.ExecuteTemplate(w, name, data)
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
