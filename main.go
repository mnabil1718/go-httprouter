package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		message, ok := error.(string) // type assertion to string
		if ok {
			fmt.Fprintf(writer, "Error: %s", message)
		}
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello world")
	})

	router.GET("/panic-example", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Intentional Server Error")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
