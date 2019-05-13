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
	fmt.Fprintf(w, "Get post(article): %v\n", vars["title"])
}

// GetPosts returns the list of all the posts
func GetPosts(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Get all posts\n")
}
