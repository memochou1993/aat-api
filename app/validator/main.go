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
	Page string `validate:"numeric"`
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
	q.Page = mutateQuery(r, "page", "1")

	return validate.Struct(q)
}
