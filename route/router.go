package route

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter will create and return the router for the app.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "From root: Hello, world!\n")
	})
	router.HandleFunc("/posts/{title}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Post: %v\n", vars["title"])
	})

	return router
}
