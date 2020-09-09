package inertia

import (
	"elipzis.com/inertia-echo/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

//
type Response struct {
	component string
	props     map[string]interface{}
	viewData  map[string]interface{}
	rootView  string
	version   string
	status    int
}

//
func NewResponse(component string, props map[string]interface{}, rootView string, version string) Response {
	var this Response
	this.component = component
	this.props = props
	this.viewData = make(map[string]interface{})
	this.rootView = rootView
	this.version = version
	this.status = http.StatusOK
	return this
}

//
func (this Response) With(key interface{}, value interface{}) Response {
	switch key.(type) {
	case string:
		this.props[key.(string)] = value
		break
	case map[string]interface{}:
		for k, v := range key.(map[string]interface{}) {
			this.props[k] = v
		}
	}
	return this
}

//
func (this Response) WithViewData(key interface{}, value interface{}) Response {
	switch key.(type) {
	case string:
		this.viewData[key.(string)] = value
		break
	case map[string]interface{}:
		for k, v := range key.(map[string]interface{}) {
			this.props[k] = v
		}

		for k, v := range key.(map[string]interface{}) {
			this.viewData[k] = v
		}
	}
	return this
}

//
func (this Response) ToResponse(c echo.Context) error {
	req := c.Request()

	var only []string
	if data := req.Header.Get(HeaderPartialData); data != "" {
		only = strings.Split(data, ",")
	}

	var props map[string]interface{}
	if only != nil && req.Header.Get(HeaderPartialData) == this.component {
		props = make(map[string]interface{})
		for _, v := range only {
			props[v] = this.props[v]
		}
	} else {
		props = this.props
	}

	// Many Question Marks here????????????
	util.WalkRecursive(props, func(prop interface{}) {
		type HandlerType func() interface{}
		switch prop.(type) {
		case func() interface{}:
			if f, ok := prop.(func() interface{}); ok {
				prop = HandlerType(f)
			}
		case Response:
			prop = prop.(Response).ToResponse(c)
		}
	})

	scheme := util.GetEnvOrDefault("SCHEME", "http")
	if req.TLS != nil || req.Header.Get("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}
	page := map[string]interface{}{
		"component": this.component,
		"props":     props,
		"url":       req.URL.String(),
		"version":   this.version,
		// Inertia-Echo-specifics
		"host":   req.Host,
		"path":   req.URL.Path,
		"scheme": scheme,
		"method": req.Method,
		"status": this.status,
	}

	if req.Header.Get(HeaderPrefix) == "true" {
		c.Response().Header().Set("Vary", "Accept")
		c.Response().Header().Set("X-Inertia", "true")
		return c.JSON(this.status, page)
	}

	this.viewData["page"] = page
	return c.Render(this.status, this.rootView, this.viewData)
}

// Allow to set a custom status
func (this Response) Status(code int) Response {
	this.status = code
	return this
}
