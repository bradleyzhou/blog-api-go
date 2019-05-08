package post

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPost returns the post by its unique title
func GetPost(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Post: %v\n", vars["title"])
}
