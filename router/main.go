package router

import (
	"github.com/gorilla/mux"
	"github.com/memochou1993/thesaurus/app/controller"
)

// NewRouter handles the routes.
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/subjects/import", controller.Import)

	return r
}
