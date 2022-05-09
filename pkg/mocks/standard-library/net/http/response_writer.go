package mocked_http

import (
	"io"
	"net/http"
	"strings"
)

type ResponseWriter interface {
	http.ResponseWriter
}

type mockedResponseWriter struct {
	header http.Header
}

func NewMockedResponseWriter() ResponseWriter {
	return &mockedResponseWriter{
		header: map[string][]string{},
	}
}

func (mrw *mockedResponseWriter) Header() http.Header {
	return mrw.header
}

func (mrw *mockedResponseWriter) Write(bytes []byte) (int, error) {
	stringReader := strings.NewReader(string(bytes))
	xb, err := io.ReadAll(stringReader)
	return len(xb), err
}

func (mrw *mockedResponseWriter) WriteHeader(statusCode int) {
	return
}
