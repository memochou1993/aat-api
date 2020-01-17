package filter

import (
	"github.com/memochou1993/thesaurus/app/mutator"
	"go.mongodb.org/mongo-driver/bson"
)

// Get gets the filters.
func Get(query *mutator.Query) bson.M {
	filters := []bson.M{}

	if query.ParentSubjectID != "" {
		filters = append(filters, bson.M{
			"parentRelationship.preferredParents.parentSubjectId": query.ParentSubjectID,
		})

		filters = append(filters, bson.M{
			"parentRelationship.nonPreferredParents.parentSubjectId": query.ParentSubjectID,
		})
	}

	if query.Term != "" {
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
	}

	if query.NoteText != "" {
		filters = append(filters, bson.M{
			"descriptiveNote.descriptiveNotes.noteText": bson.M{
				"$regex": ".*" + query.Term + ".*",
			},
		})
	}

	return bson.M{
		"$or": filters,
	}
}
