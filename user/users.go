package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetUser returns the user by user name
func GetUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Get user: %v\n", vars["name"])
}
