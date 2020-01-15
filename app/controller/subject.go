package controller

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload" // initialize
	"github.com/memochou1993/thesaurus/app/formatter"
	"github.com/memochou1993/thesaurus/app/model"
	"github.com/memochou1993/thesaurus/app/mutator"
	"github.com/memochou1993/thesaurus/app/parser"
	"github.com/memochou1993/thesaurus/app/validator"
	"go.mongodb.org/mongo-driver/bson"
)

// GetSubjects gets the subjects.
func GetSubjects(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	model := model.Subjects{}
	validator := validator.Query{}
	mutator := mutator.Query{}
	formatter := formatter.Payload{}

	if err := validator.Validate(r); err != nil {
		response(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := mutator.Mutate(&validator); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := model.FindAll(&mutator); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	formatter.Set(&model.Subjects)

	response(w, http.StatusOK, formatter)
}

// GetSubject gets the subject.
func GetSubject(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	model := model.Subject{}
	formatter := formatter.Payload{}

	vars := mux.Vars(r)

	if err := model.Find(bson.M{"subjectId": vars["id"]}); err != nil {
		response(w, http.StatusNotFound, nil)
		return
	}

	formatter.Set(&model)

	response(w, http.StatusOK, formatter)
}

// ImportSubjects imports the subjects.
func ImportSubjects(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	model := model.Subjects{}

	file := os.Getenv("RESOURCE_PATH")
	parser.Parse(file, &model)

	if err := model.BulkUpsert(); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := model.PopulateIndex(); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusCreated, nil)
}
