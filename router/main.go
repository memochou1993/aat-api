package router

import (
	"github.com/gorilla/mux"
	"github.com/memochou1993/thesaurus/app/controller"
)

// NewRouter handles the routes.
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/subjects", controller.GetSubjects).Methods("GET")
	api.HandleFunc("/subjects/{id}", controller.GetSubject).Methods("GET")
	api.HandleFunc("/subjects", controller.ImportSubjects).Methods("PUT")

	return r
}
