package route

import (
	"github.com/bradleyzhou/blog-api-go/user"
	"github.com/gorilla/mux"
)

func registerUsersRoutes(router *mux.Router) {
	r := router.PathPrefix("/users").Subrouter()

	r.HandleFunc("/{name}", user.GetUser)
}
