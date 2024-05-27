package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

//go:embed resources/1.pdf
var pdfContent string

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	dir, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}
	router.ServeFiles("/static/*filepath", http.FS(dir))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/static/1.pdf", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, pdfContent, string(body))

}
