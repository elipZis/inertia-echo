package inertia

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/elipzis/inertia-echo/util"
	_ "github.com/joho/godotenv/autoload"
	"io/ioutil"
)

// The Base "X-Inertia" header prefix
const HeaderPrefix = "X-Inertia"
const HeaderVersion = HeaderPrefix + "-Version"
const HeaderLocation = HeaderPrefix + "-Location"
const HeaderPartialData = HeaderPrefix + "-Partial-Data"

// Get the configured root view from the environment
var RootView = util.GetEnvOrDefault("INERTIA_ROOT_VIEW", "app.html")

// Get the configured public path from the environment
var PublicPath = util.GetEnvOrDefault("PUBLIC_PATH", "public")

//
type Inertia struct {
	sharedProps map[string]interface{}
	version     interface{}
}

// Instance a new Inertia Handler
func NewInertia() (this *Inertia) {
	this = new(Inertia)
	this.sharedProps = make(map[string]interface{})
	return this
}

// Render the given component with props and version
func (this *Inertia) Render(component string, props map[string]interface{}) Response {
	return NewResponse(component, util.MergeMaps(this.sharedProps, props), RootView, this.GetVersion())
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
func (this *Inertia) SetMixVersion() bool {
	fileData, err := ioutil.ReadFile(PublicPath + "/mix-manifest.json")
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
