package controller

import (
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload" // initialize
	"github.com/memochou1993/thesaurus/app/model"
	"github.com/memochou1993/thesaurus/app/mutator"
	"github.com/memochou1993/thesaurus/app/parser"
	"github.com/memochou1993/thesaurus/app/validator"
)

var (
	vocabulary     model.Vocabulary
	queryValidator validator.Query
	queryMutator   mutator.Query
)

func response(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Index displays a listing of the resource.
func Index(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if err := queryValidator.Validate(r); err != nil {
		response(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := queryMutator.Mutate(&queryValidator); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := vocabulary.FindAll(&queryMutator); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, vocabulary.Subjects)
}

// Import imports the resource from a XML file.
func Import(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	file := os.Getenv("RESOURCE_PATH")
	parser.Parse(file, &vocabulary)

	if err := vocabulary.BulkUpsert(); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := vocabulary.PopulateIndex(); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusCreated, nil)
}
