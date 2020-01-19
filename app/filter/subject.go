package filter

import (
	"github.com/memochou1993/thesaurus/app/mutator"
	"go.mongodb.org/mongo-driver/bson"
)

// Subject struct
type Subject struct {
	Fliter bson.M
}

// Set sets the filter.
func (s *Subject) Set(query *mutator.Query) {
	filters := []bson.M{}

	filters = append(filters, bson.M{
		"parentRelationship.preferredParents.parentSubjectId": query.ParentSubjectID,
	})

	filters = append(filters, bson.M{
		"parentRelationship.nonPreferredParents.parentSubjectId": query.ParentSubjectID,
	})

	filters = append(filters, bson.M{
		"term.preferredTerms.termText": bson.M{
			"$regex": ".*" + query.Term + ".*",
		},
	})

	filters = append(filters, bson.M{
		"term.nonPreferredTerms.termText": bson.M{
			"$regex": ".*" + query.Term + ".*",
		},
	})

	filters = append(filters, bson.M{
		"descriptiveNote.descriptiveNotes.noteText": bson.M{
			"$regex": ".*" + query.Term + ".*",
		},
	})

	s.Fliter = bson.M{
		"$or": filters,
	}
}
