package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/globalsign/mgo/bson"
	_ "github.com/joho/godotenv/autoload"
	"github.com/memochou1993/thesaurus/app/model"
	"github.com/memochou1993/thesaurus/app/parser"
)

var (
	vocabulary = model.Vocabulary{}
	subject    = model.Subject{}
)

func response(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Import imports subjects from a XML file.
func Import(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	resource := os.Getenv("RESOURCE_PATH")

	parser.Parse(resource, &vocabulary)

	for _, subject := range vocabulary.Subjects {
		subject.Upsert(bson.M{"subjectId": subject.SubjectID}, subject)
	}

	response(w, http.StatusCreated, nil)
}
