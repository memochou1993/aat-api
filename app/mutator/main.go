package mutator

import (
	"strconv"

	"github.com/memochou1993/thesaurus/app/validator"
)

// Query struct
type Query struct {
	ParentSubjectID string
	Term            string
	Page            int64
	PageSize        int64
}

func mutatePage(v string) (int64, error) {
	i, err := strconv.ParseInt(v, 10, 64)

	if i < 1 {
		i = 1
	}

	return i, err
}

func mutatePageSize(v string) (int64, error) {
	i, err := strconv.ParseInt(v, 10, 64)

	if i < 1 {
		i = 10
	}

	return i, err
}

// Mutate mutates the query.
func (q *Query) Mutate(query *validator.Query) error {
	page, err := mutatePage(query.Page)
	pageSize, err := mutatePageSize(query.PageSize)

	q.ParentSubjectID = query.ParentSubjectID
	q.Term = query.Term
	q.Page = page
	q.PageSize = pageSize

	return err
}
