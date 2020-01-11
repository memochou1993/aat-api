package mutator

import (
	"github.com/memochou1993/thesaurus/app/validator"
	"strconv"
)

// Query struct
type Query struct {
	Page int64
}

func mutatePage(value string) (int64, error) {
	page, err := strconv.ParseInt(value, 10, 64)

	if page < 1 {
		page = 1
	}

	return page, err
}

// Mutate mutates the query.
func (q *Query) Mutate(query *validator.Query) error {
	page, err := mutatePage(query.Page)

	q.Page = page

	return err
}
