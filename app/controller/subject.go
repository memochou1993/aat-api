package controller

import (
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload" // initialize
	"github.com/memochou1993/thesaurus/app/model"
	"github.com/memochou1993/thesaurus/app/parser"
	"go.mongodb.org/mongo-driver/bson"
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

	file := os.Getenv("RESOURCE_PATH")
	parser.Parse(file, &vocabulary)

	for _, subject := range vocabulary.Subjects {
		query := bson.M{"subjectId": subject.SubjectID}
		update := bson.M{"$set": subject}

		if err := subject.Upsert(query, update); err != nil {
			response(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	response(w, http.StatusCreated, nil)
}
