package controller

import (
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload" // initialize
	"github.com/memochou1993/thesaurus/app/formatter"
	"github.com/memochou1993/thesaurus/app/model"
	"github.com/memochou1993/thesaurus/app/mutator"
	"github.com/memochou1993/thesaurus/app/parser"
	"github.com/memochou1993/thesaurus/app/validator"
)

var (
	subjectModel       model.Subjects
	queryValidator     validator.Query
	queryMutator       mutator.Query
	payloadTransformer formatter.Payload
)

// GetSubjects gets the subjects.
func GetSubjects(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if err := queryValidator.Validate(r); err != nil {
		response(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := queryMutator.Mutate(&queryValidator); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := subjectModel.FindAll(&queryMutator); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	payloadTransformer.Get(&subjectModel)

	response(w, http.StatusOK, payloadTransformer)
}

// ImportSubjects imports the subjects.
func ImportSubjects(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	file := os.Getenv("RESOURCE_PATH")
	parser.Parse(file, &subjectModel)

	if err := subjectModel.BulkUpsert(); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := subjectModel.PopulateIndex(); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusCreated, nil)
}
