package controller

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/globalsign/mgo/bson"
	"github.com/memochou1993/thesaurus/app/model"
)

const (
	resource = "./storage/vocabulary.xml"
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

func parse() {
	file, err := os.Open(resource)
	defer file.Close()

	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return
	}

	xml.Unmarshal(data, &vocabulary)
}

// Import imports subjects from XML file.
func Import(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	parse()

	for i := 0; i < len(vocabulary.Subjects); i++ {
		subject.Upsert(bson.M{"subjectId": vocabulary.Subjects[i].SubjectID}, vocabulary.Subjects[i])
	}

	response(w, http.StatusCreated, nil)
}
