package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	BootstrapHttpServer()
}

func BootstrapHttpServer() {
	router := mux.NewRouter().PathPrefix("/v1").Subrouter()
	RegisterAccountEndpoints(router)
	RegisterSwaggerEndpoints(router)
	RegisterGamesEndpoints(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("Starting the server at address " + srv.Addr)
	var err = srv.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start http server %s", err.Error())
	}
}
