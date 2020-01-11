package mutator

import (
	"github.com/memochou1993/thesaurus/app/validator"
	"strconv"
)

// Query struct
type Query struct {
	Page int64
}

// Mutate mutates the query.
func (q *Query) Mutate(query *validator.Query) (*Query, error) {
	page, err := strconv.ParseInt(query.Page, 10, 64)

	q.Page = page

	return q, err
}
