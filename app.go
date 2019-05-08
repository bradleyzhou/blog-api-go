package blog

import (
	"log"
	"net/http"
	"time"

	"github.com/bradleyzhou/blog-api-go/route"
)

func main() {
	router := route.NewRouter()

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
