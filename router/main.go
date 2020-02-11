package router

import (
	"github.com/gorilla/mux"
	"github.com/memochou1993/thesaurus/app/controller"
)

// NewRouter handles the routes.
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/subjects", controller.GetSubjects).Methods("GET")
	r.HandleFunc("/subjects/{id}", controller.GetSubject).Methods("GET")
	r.HandleFunc("/subjects", controller.ImportSubjects).Methods("PUT")

	return r
}
