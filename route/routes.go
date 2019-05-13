package route

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterV1Routes composed all the /api/v1 routes
func RegisterV1Routes(router *mux.Router) {
	v1r := router.PathPrefix("/api/v1").Subrouter()

	v1r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "From root: Hello, world!\n")
	})

	registerPostsRoutes(v1r)
	registerUsersRoutes(v1r)
}
