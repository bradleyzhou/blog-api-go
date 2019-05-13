package route

import (
	"github.com/bradleyzhou/blog-api-go/post"
	"github.com/gorilla/mux"
)

func registerPostsRoutes(router *mux.Router) {
	r := router.PathPrefix("/posts").Subrouter()

	r.HandleFunc("", post.GetPosts)
	r.HandleFunc("/", post.GetPosts)
	r.HandleFunc("/{title}", post.GetPost)
}
