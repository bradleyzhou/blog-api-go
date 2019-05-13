package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bradleyzhou/blog-api-go/route"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	route.RegisterV1Routes(router)

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
