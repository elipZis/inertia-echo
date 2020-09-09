package util

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

// https://github.com/tightenco/ziggy base struct to export compatible Echo routes
type Ziggy struct {
	BaseDomain   string                `json:"baseDomain"`
	BasePort     int                   `json:"basePort"`
	BaseProtocol string                `json:"baseProtocol"`
	BaseUrl      string                `json:"baseUrl"`
	Group        string                `json:"group"`
	Routes       map[string]ZiggyRoute `json:"namedRoutes"`
}

// A single https://github.com/tightenco/ziggy route
type ZiggyRoute struct {
	Uri     string   `json:"uri"`
	Methods []string `json:"methods"`
	Domain  string   `json:"domain"`
}

//
func NewZiggy(echo *echo.Echo, page map[string]interface{}) Ziggy {
	var this Ziggy

	this.BaseProtocol = "http"
	var s []string
	if host, ok := page["host"]; ok {
		s = strings.Split(host.(string), ":")
	} else {
		s = strings.Split(echo.Server.Addr, "::")
	}
	this.BaseDomain = s[0]
	this.BasePort = 80
	if len(s) > 1 {
		port, _ := strconv.Atoi(s[1])
		this.BasePort = port
	}
	this.BaseUrl = this.BaseProtocol + "://" + this.BaseDomain + ":" + strconv.Itoa(this.BasePort) + "/"

	this.Routes = make(map[string]ZiggyRoute, len(echo.Routes()))
	for _, route := range echo.Routes() {
		if ziggyRoute, ok := this.Routes[route.Name]; ok {
			ziggyRoute.Methods = append(ziggyRoute.Methods, route.Method)
			this.Routes[route.Name] = ziggyRoute
		} else {
			this.Routes[route.Name] = ZiggyRoute{
				Uri: route.Path,
				Methods: []string{
					route.Method,
				},
			}
		}
	}

	return this
}
