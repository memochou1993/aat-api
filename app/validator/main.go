package validator

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

// Query struct
type Query struct {
	ParentSubjectID string `validate:""`
	Term            string `validate:""`
	NoteText        string `validate:""`
	Page            string `validate:"numeric"`
	PageSize        string `validate:"numeric"`
}

func init() {
	validate = validator.New()
}

func get(r *http.Request, key string, defaultValue string) string {
	query := r.URL.Query().Get(key)

	if query == "" {
		query = defaultValue
	}

	return query
}

// Validate validates the query.
func (q *Query) Validate(r *http.Request) error {
	q.ParentSubjectID = get(r, "parentSubjectId", "")
	q.Term = get(r, "term", "")
	q.NoteText = get(r, "noteText", "")
	q.Page = get(r, "page", "1")
	q.PageSize = get(r, "pageSize", "10")

	return validate.Struct(q)
}
