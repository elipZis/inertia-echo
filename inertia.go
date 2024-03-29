package inertia

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/elipzis/inertia-echo/util"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"sync"
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

	sharedProps      map[string]map[string]interface{}
	sharedPropsMutex *sync.Mutex
	version          interface{}
}

//
type InertiaConfig struct {
	Echo *echo.Echo

	PublicPath       string
	TemplatesPath    string
	RootView         string
	TemplateFuncMap  template.FuncMap
	HTTPErrorHandler echo.HTTPErrorHandler
	RequestIDConfig  middleware.RequestIDConfig
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

	this.RequestIDConfig = middleware.DefaultRequestIDConfig

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
	this.sharedProps = make(map[string]map[string]interface{})
	this.sharedPropsMutex = &sync.Mutex{}
	this.config.Echo.Renderer = this
	this.config.Echo.HTTPErrorHandler = this.config.HTTPErrorHandler
	log.Printf("[Inertia] Loading templates out of %s", this.config.TemplatesPath)
	this.templates = template.Must(template.New("").Funcs(this.config.TemplateFuncMap).ParseGlob(this.config.TemplatesPath))
	// Try to set a version off of the mix-manifest, if any
	this.SetMixVersion()
	// Register a unique id generator to identify requests
	this.config.Echo.Use(middleware.RequestIDWithConfig(this.config.RequestIDConfig))

	return this
}

// Render renders a template document
func (this *Inertia) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Always empty the shared props for this request
	sharedProps := this.Shared(c)
	isMap := reflect.TypeOf(data).Kind() == reflect.Map
	if this.templates.Lookup(name) != nil {
		if isMap {
			viewContext := data.(map[string]interface{})
			viewContext["reverse"] = c.Echo().Reverse
			viewContext["shared"] = sharedProps
		}
		return this.templates.ExecuteTemplate(w, name, data)
	}

	if isMap {
		if sharedProps != nil {
			return NewResponse(name, util.MergeMaps(sharedProps, data.(map[string]interface{})), this.config.RootView, this.GetVersion()).Status(c.Response().Status).ToResponse(c)
		} else {
			return NewResponse(name, data.(map[string]interface{}), this.config.RootView, this.GetVersion()).Status(c.Response().Status).ToResponse(c)
		}
	}
	return NewResponse(name, sharedProps, this.config.RootView, this.GetVersion()).Status(c.Response().Status).ToResponse(c)
}

// Share a key/value pairs with every response
func (this *Inertia) Share(c echo.Context, key string, value interface{}) {
	rid := c.Request().Header.Get(echo.HeaderXRequestID)
	this.sharedPropsMutex.Lock()
	if reqSharedProps, ok := this.sharedProps[rid]; ok {
		reqSharedProps[key] = value
	} else {
		this.sharedProps[rid] = map[string]interface{}{
			key: value,
		}
	}
	this.sharedPropsMutex.Unlock()
}

// Share multiple key/values with every response
func (this *Inertia) Shares(c echo.Context, values map[string]interface{}) {
	rid := c.Request().Header.Get(echo.HeaderXRequestID)
	this.sharedPropsMutex.Lock()
	if _, ok := this.sharedProps[rid]; !ok {
		this.sharedProps[rid] = make(map[string]interface{})
	}
	this.sharedPropsMutex.Unlock()

	for key, value := range values {
		this.sharedPropsMutex.Lock()
		this.sharedProps[rid][key] = value
		this.sharedPropsMutex.Unlock()
	}
}

// Get a specific key-value from the shared information
func (this *Inertia) GetShared(c echo.Context, key string) (interface{}, bool) {
	rid := c.Request().Header.Get(echo.HeaderXRequestID)
	this.sharedPropsMutex.Lock()
	if reqSharedProps, ok := this.sharedProps[rid]; ok {
		value, ok := reqSharedProps[key]
		this.sharedPropsMutex.Unlock()
		return value, ok
	}
	return nil, false
}

// Returns the shared props (if any) and deletes them
func (this *Inertia) Shared(c echo.Context) map[string]interface{} {
	rid := c.Request().Header.Get(echo.HeaderXRequestID)
	this.sharedPropsMutex.Lock()
	sharedProps := this.sharedProps[rid]
	delete(this.sharedProps, rid)
	this.sharedPropsMutex.Unlock()
	return sharedProps
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
