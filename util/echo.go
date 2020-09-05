package util

import (
	"github.com/labstack/echo/v4"
	"strings"
)

//
func GetBaseUrl(c echo.Context) string {
	req, scheme := c.Request(), c.Scheme()
	host := req.Host
	url := scheme + "://" + host + req.RequestURI
	return strings.TrimSuffix(url, "/")
}
