package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	stringResult := "Product 1"
	router := httprouter.New()
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		productId := params.ByName("id")
		fmt.Fprint(writer, "Product "+productId)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, stringResult, string(body))

}

func TestParamsCatchAll(t *testing.T) {
	stringResult := "/src/images/image.jpg"
	router := httprouter.New()
	router.GET("/static/*imgpath", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		imgpath := params.ByName("imgpath")
		fmt.Fprint(writer, imgpath)
	})

	// first / after static is included as catch all params
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/static/src/images/image.jpg", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, stringResult, string(body))
}
