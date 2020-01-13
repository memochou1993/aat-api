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
	SubjectID string `validate:""`
	Term      string `validate:""`
	Page      string `validate:"numeric"`
	PageSize  string `validate:"numeric"`
}

func init() {
	validate = validator.New()
}

func mutateQuery(r *http.Request, key string, defaultValue string) string {
	query := r.URL.Query().Get(key)

	if query == "" {
		query = defaultValue
	}

	return query
}

// Validate validates the query.
func (q *Query) Validate(r *http.Request) error {
	q.SubjectID = mutateQuery(r, "subjectId", "")
	q.Term = mutateQuery(r, "term", "")
	q.Page = mutateQuery(r, "page", "1")
	q.PageSize = mutateQuery(r, "pageSize", "10")

	return validate.Struct(q)
}
