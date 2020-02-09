package controller

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload" // initialize
	"github.com/memochou1993/thesaurus/app/formatter"
	"github.com/memochou1993/thesaurus/app/model"
	"github.com/memochou1993/thesaurus/app/mutator"
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

	if len(model.Subjects) == 0 {
		formatter.Set([]string{})
	}

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

	resource := os.Getenv("RESOURCE_PATH")
	file, err := os.Open(resource)
	defer file.Close()

	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	decoder := xml.NewDecoder(file)

	for {
		token, err := decoder.Token()

		if err == io.EOF {
			break
		}

		if err != nil {
			response(w, http.StatusInternalServerError, err.Error())
			return
		}

		if token == nil {
			break
		}

		switch element := token.(type) {
		case xml.StartElement:
			switch element.Name.Local {
			case "Subject":
				model := model.Subject{}

				if err = decoder.DecodeElement(&model, &element); err != nil {
					response(w, http.StatusInternalServerError, err.Error())
					return
				}

				if err := model.Upsert(); err != nil {
					response(w, http.StatusInternalServerError, err.Error())
					return
				}
			}
		}
	}

	model := model.Subjects{}

	if err := model.PopulateIndex(); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusCreated, nil)
}
