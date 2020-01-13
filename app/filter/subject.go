package filter

import (
	"github.com/memochou1993/thesaurus/app/mutator"
	"go.mongodb.org/mongo-driver/bson"
)

// Get gets the filter.
func Get(query *mutator.Query) bson.M {
	if query.SubjectID != "" {
		return bson.M{
			"subjectId": query.SubjectID,
		}
	}

	if query.Term != "" {
		return bson.M{
			"$or": []bson.M{
				bson.M{
					"descriptiveNote.descriptiveNotes.noteText": bson.M{
						"$regex": ".*" + query.Term + ".*",
					},
				},
				bson.M{
					"term.preferredTerms.termText": bson.M{
						"$regex": ".*" + query.Term + ".*",
					},
				},
				bson.M{
					"term.nonPreferredTerms.termText": bson.M{
						"$regex": ".*" + query.Term + ".*",
					},
				},
			},
		}
	}

	return bson.M{}
}
