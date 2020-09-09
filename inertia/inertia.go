package inertia

import (
	"crypto/md5"
	"elipzis.com/inertia-echo/util"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

// The Base "X-Inertia" header prefixes
const HeaderPrefix = "X-Inertia"
const HeaderVersion = HeaderPrefix + "-Version"
const HeaderLocation = HeaderPrefix + "-Location"
const HeaderPartialData = HeaderPrefix + "-Partial-Data"

//
type Inertia struct {
	config InertiaConfig

	templates *template.Template

	sharedProps map[string]interface{}
	version     interface{}
}

//
type InertiaConfig struct {
	Echo *echo.Echo

	PublicPath       string
	TemplatesPath    string
	RootView         string
	TemplateFuncMap  template.FuncMap
	HTTPErrorHandler echo.HTTPErrorHandler
}

//
var DefaultHTTPErrorHandler = func(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	_ = c.Render(code, "Error", map[string]interface{}{
		"status": code,
	})
}

// Create a default Inertia Config to use without the hassle of setting up everything
func NewDefaultInertiaConfig(e *echo.Echo) (this InertiaConfig) {
	this = InertiaConfig{}
	this.Echo = e

	// Get the configured root view from the environment
	this.RootView = util.GetEnvOrDefault("INERTIA_ROOT_VIEW", "app.html")
	// Get the configured public path from the environment
	this.PublicPath = util.GetEnvOrDefault("INERTIA_PUBLIC_PATH", "public")
	// Get the configured templates path from the environment
	this.TemplatesPath = util.GetEnvOrDefault("INERTIA_RESOURCES_PATH", "resources") + "/views/*.html"

	// Set a default error handler to render a default error page
	this.HTTPErrorHandler = DefaultHTTPErrorHandler

	// Register convenient template functions
	this.TemplateFuncMap = template.FuncMap{
		"inertia":         util.Inertia,
		"json_encode":     util.JsonEncode,
		"json_encode_raw": util.JsonEncodeRaw,
		"mix":             util.Mix,
		"routes": func() template.JS {
			retVal, _ := json.Marshal(e.Routes())
			return template.JS(string(retVal))
		},
		"routes_ziggy": func(v interface{}) template.HTML {
			ziggy := util.NewZiggy(e, v.(map[string]interface{}))
			retVal, _ := json.Marshal(ziggy)
			return template.HTML(fmt.Sprintf("<script>var Ziggy = %s;</script>", string(retVal)))
		},
	}

	return this
}

// Instance a new Inertia Handler with a default config
func NewInertia(echo *echo.Echo) (this *Inertia) {
	return NewInertiaWithConfig(NewDefaultInertiaConfig(echo))
}

// Instance a new Inertia Handler with a configuration
func NewInertiaWithConfig(config InertiaConfig) (this *Inertia) {
	if config.Echo == nil {
		log.Fatal("[Inertia] echo.Echo reference required in the given config!")
	}

	this = new(Inertia)
	this.config = config
	this.sharedProps = make(map[string]interface{})
	this.config.Echo.Renderer = this
	this.config.Echo.HTTPErrorHandler = this.config.HTTPErrorHandler
	this.templates = template.Must(template.New("").Funcs(this.config.TemplateFuncMap).ParseGlob(this.config.TemplatesPath))
	// Try to set a version off of the mix-manifest, if any
	this.SetMixVersion()

	return this
}

// Render renders a template document
func (this *Inertia) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := this.templates.ExecuteTemplate(w, name, data)
	log.Print(err)
	return err
	isMap := reflect.TypeOf(data).Kind() == reflect.Map
	if this.templates.Lookup(name) != nil {
		if isMap {
			viewContext := data.(map[string]interface{})
			viewContext["reverse"] = c.Echo().Reverse
			viewContext["shared"] = this.sharedProps
		}
		return this.templates.ExecuteTemplate(w, name, data)
	}

	if isMap {
		return NewResponse(name, util.MergeMaps(this.sharedProps, data.(map[string]interface{})), this.config.RootView, this.GetVersion()).Status(c.Response().Status).ToResponse(c)
	}
	return NewResponse(name, this.sharedProps, this.config.RootView, this.GetVersion()).Status(c.Response().Status).ToResponse(c)
}

// Share a key/value pairs with every response
func (this *Inertia) Share(key string, value interface{}) {
	this.sharedProps[key] = value
}

// Share multiple key/values with every response
func (this *Inertia) Shares(values map[string]interface{}) {
	for key, value := range values {
		this.sharedProps[key] = value
	}
}

// Get a specific key-value from the shared information
func (this *Inertia) GetShared(key string) (interface{}, bool) {
	value, ok := this.sharedProps[key]
	return value, ok
}

func (this *Inertia) Shared() map[string]interface{} {
	return this.sharedProps
}

// Set a version callback "func() string"
func (this *Inertia) Version(version func() string) {
	this.version = version
}

// Set a version string
func (this *Inertia) SetVersion(version string) {
	this.version = version
}

// Create a version hash off of the mix-manifest.json file md5
func (this *Inertia) SetMixVersion(mixManifestPath ...string) bool {
	filePath := this.config.PublicPath + "/mix-manifest.json"
	if len(mixManifestPath) > 0 {
		filePath = mixManifestPath[0]
	}
	fileData, err := ioutil.ReadFile(filePath)
	if err == nil {
		hash := md5.New()
		hash.Write(fileData)
		this.version = hex.EncodeToString(hash.Sum(nil))
		return true
	}
	return false
}

//
func (this *Inertia) GetVersion() string {
	if this.version != nil {
		type HandlerType func() string
		switch this.version.(type) {
		case func() string:
			if f, ok := this.version.(func() string); ok {
				this.version = HandlerType(f)
			}
			break
		}

		return this.version.(string)
	}

	return ""
}
