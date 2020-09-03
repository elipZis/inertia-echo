package inertia

import (
	"bufio"
	"bytes"
	"elipzis.com/inertia-echo/service"
	"fmt"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net"
	"net/http"
	"strconv"
)

//
type MiddlewareConfig struct {
	Inertia *Inertia
	Skipper middleware.Skipper
}

type DumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *DumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *DumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *DumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *DumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

// The Inertia Middleware to check every request for what it needs
func MiddlewareWithConfig(config MiddlewareConfig) echo.MiddlewareFunc {
	if config.Inertia == nil {
		config.Inertia = NewInertia()
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip, if configured and true
			if config.Skipper != nil && config.Skipper(c) {
				return next(c)
			}

			// Response
			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(c.Response().Writer, resBody)
			writer := &DumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
			c.Response().Writer = writer

			// Run Inertia post
			if err := next(c); err != nil {
				c.Error(err)

				fmt.Println(err)

				code := http.StatusInternalServerError
				message := err.Error()
				if he, ok := err.(*echo.HTTPError); ok {
					code = he.Code
					message = he.Message.(string)
				}
				errorMsg := map[string]interface{}{
					strconv.Itoa(code): message,
				}

				// Add general errors in case some pop up
				errors, ok := config.Inertia.GetShared("errors")
				if !ok {
					config.Inertia.Share("errors", errorMsg)
				} else {
					config.Inertia.Share("errors", service.MergeMaps(errors.(map[string]interface{}), errorMsg))
				}
			}

			req := c.Request()
			res := c.Response()

			if req.Header.Get(HeaderPrefix) == "" {
				return nil
			}

			if req.Method == "GET" && req.Header.Get(HeaderVersion) != config.Inertia.GetVersion() {
				// $request->session()->reflash();???
				if s, err := session.Get("session", c); err == nil {
					flashes := s.Flashes()
					config.Inertia.Share("flash", flashes)
					for _, flash := range flashes {
						s.AddFlash(flash)
					}
				}

				res.Header().Set(HeaderLocation, req.URL.String())
				return c.String(http.StatusConflict, "")
			}

			if exists, _ := service.InArray(req.Method, []string{"PUT", "PATCH", "DELETE"}); exists && res.Status == 302 {
				res.Status = http.StatusSeeOther
			}

			// Error handling
			fmt.Println(res.Status)
			fmt.Println(c)
			fmt.Println(res.Committed)
			fmt.Println(string(resBody.Bytes()))

			return nil
		}
	}
}
