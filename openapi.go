package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterSwaggerEndpoints(router *mux.Router) {
	router.HandleFunc("/openapi.yaml", YamlHandler).Methods(http.MethodGet)
}

func YamlHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(writer, request, "openapi/openapi.yaml")
}
