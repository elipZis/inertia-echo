package util

import (
	"bufio"
	"io"
	"net"
	"net/http"
)

// An intermediate ResponseWriter to dump content
type Dumper struct {
	io.Writer
	http.ResponseWriter
}

//
func (w *Dumper) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

//
func (w *Dumper) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

//
func (w *Dumper) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

//
func (w *Dumper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
