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
	Page string `validate:"required,numeric"`
}

func init() {
	validate = validator.New()
}

// Validate func
func (q *Query) Validate(r *http.Request) (*Query, error) {
	query := r.URL.Query()

	q = &Query{
		Page: query.Get("page"),
	}

	return q, validate.Struct(q)
}
