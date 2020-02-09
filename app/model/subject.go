package model

import (
	"context"
	"time"

	"github.com/memochou1993/thesaurus/app/filter"
	"github.com/memochou1993/thesaurus/app/mutator"
	"github.com/memochou1993/thesaurus/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collection = "subjects"
)

// Subjects struct
type Subjects struct {
	Subjects []Subject `xml:"Subject" bson:"subjects" json:"subjects"`
}

// Subject struct
type Subject struct {
	SubjectID           string `xml:"Subject_ID,attr" bson:"subjectId" json:"subjectId"`
	ParentRelationships struct {
		PreferredParent    []ParentRelationship `xml:"Preferred_Parent" bson:"preferredParents" json:"preferredParents"`
		NonPreferredParent []ParentRelationship `xml:"Non-Preferred_Parent" bson:"nonPreferredParents" json:"nonPreferredParents"`
	} `xml:"Parent_Relationships" bson:"parentRelationship" json:"parentRelationship"`
	DescriptiveNotes struct {
		DescriptiveNote []struct {
			NoteText         string `xml:"Note_Text" bson:"noteText" json:"noteText"`
			NoteLanguage     string `xml:"Note_Language" bson:"noteLanguage" json:"noteLanguage"`
			NoteContributors struct {
				NoteContributor []struct {
					ContributorID string `xml:"Contributor_id" bson:"contributorId" json:"contributorId"`
				} `xml:"Note_Contributor" bson:"noteContributors" json:"noteContributors"`
			} `xml:"Note_Contributors" bson:"noteContributor" json:"noteContributor"`
			NoteSources struct {
				NoteSource []struct {
					Source Source `xml:"Source" bson:"source" json:"source"`
				} `xml:"Note_Source" bson:"noteSources" json:"noteSources"`
			} `xml:"Note_Sources" bson:"noteSource" json:"noteSource"`
		} `xml:"Descriptive_Note" bson:"descriptiveNotes" json:"descriptiveNotes"`
	} `xml:"Descriptive_Notes" bson:"descriptiveNote" json:"descriptiveNote"`
	RecordType   string `xml:"Record_Type" bson:"recordType" json:"recordType"`
	MergedStatus string `xml:"Merged_Status" bson:"mergedStatus" json:"mergedStatus"`
	Hierarchy    string `xml:"Hierarchy" bson:"hierarchy" json:"hierarchy"`
	SortOrder    string `xml:"Sort_Order" bson:"sortOrder" json:"sortOrder"`
	Terms        struct {
		PreferredTerm    []Term `xml:"Preferred_Term" bson:"preferredTerms" json:"preferredTerms"`
		NonPreferredTerm []Term `xml:"Non-Preferred_Term" bson:"nonPreferredTerms" json:"nonPreferredTerms"`
	} `xml:"Terms" bson:"term" json:"term"`
	AssociativeRelationships struct {
		AssociativeRelationship []AssociativeRelationship `xml:"Associative_Relationship" bson:"associativeRelationships" json:"associativeRelationships"`
	} `xml:"Associative_Relationships" bson:"associativeRelationship" json:"associativeRelationship"`
	SubjectContributors struct {
		SubjectContributor []struct {
			ContributorID string `xml:"Contributor_id" bson:"contributorId" json:"contributorId"`
		} `xml:"Subject_Contributor" bson:"subjectContributors" json:"subjectContributors"`
	} `xml:"Subject_Contributors" bson:"subjectContributor" json:"subjectContributor"`
	SubjectSources struct {
		SubjectSource []struct {
			Source Source `xml:"Source" bson:"source" json:"source"`
		} `xml:"Subject_Source" bson:"subjectSources" json:"subjectSources"`
	} `xml:"Subject_Sources" bson:"subjectSource" json:"subjectSource"`
}

// ParentRelationship struct
type ParentRelationship struct {
	ParentSubjectID  string `xml:"Parent_Subject_ID" bson:"parentSubjectId" json:"parentSubjectId"`
	RelationshipType string `xml:"Relationship_Type" bson:"relationshipType" json:"relationshipType"`
	HistoricFlag     string `xml:"Historic_Flag" bson:"historicFlag" json:"historicFlag"`
	ParentString     string `xml:"Parent_String" bson:"parentString" json:"parentString"`
	HierRelType      string `xml:"Hier_Rel_Type" bson:"hierRelType" json:"hierRelType"`
}

// AssociativeRelationship struct
type AssociativeRelationship struct {
	RelationshipType string `xml:"Relationship_Type" bson:"relationshipType" json:"relationshipType"`
	RelatedSubjectID struct {
		VPSubjectID string `xml:"VP_Subject_ID" bson:"vpSubjectId" json:"vpSubjectId"`
	} `xml:"Related_Subject_ID" bson:"relatedSubjectId" json:"relatedSubjectId"`
	HistoricFlag string `xml:"Historic_Flag" bson:"historicFlag" json:"historicFlag"`
}

// Term struct
type Term struct {
	TermText      string `xml:"Term_Text" bson:"termText" json:"termText"`
	DisplayName   string `xml:"Display_Name" bson:"displayName" json:"displayName"`
	HistoricFlag  string `xml:"Historic_Flag" bson:"historicFlag" json:"historicFlag"`
	Vernacular    string `xml:"Vernacular" bson:"vernacular" json:"vernacular"`
	TermID        string `xml:"Term_ID" bson:"termId" json:"termId"`
	TermLanguages struct {
		TermLanguage []struct {
			Language     string `xml:"Language" bson:"language" json:"language"`
			Preferred    string `xml:"Preferred" bson:"preferred" json:"preferred"`
			Qualifier    string `xml:"Qualifier" bson:"qualifier" json:"qualifier"`
			TermType     string `xml:"Term_Type" bson:"termType" json:"termType"`
			PartOfSpeech string `xml:"Part_of_Speech" bson:"partOfSpeech" json:"partOfSpeech"`
			LangStat     string `xml:"Lang_Stat" bson:"langStat" json:"langStat"`
		} `xml:"Term_Language" bson:"termLanguages" json:"termLanguages"`
	} `xml:"Term_Languages" bson:"termLanguage" json:"termLanguage"`
	TermContributors struct {
		TermContributor []struct {
			ContributorID string `xml:"Contributor_id" bson:"contributorId" json:"contributorId"`
			Preferred     string `xml:"Preferred" bson:"preferred" json:"preferred"`
		} `xml:"Term_Contributor" bson:"termContributors" json:"termContributors"`
	} `xml:"Term_Contributors" bson:"termContributor" json:"termContributor"`
	TermSources struct {
		TermSource []struct {
			Source    Source `xml:"Source" bson:"source" json:"source"`
			Page      string `xml:"Page" bson:"page" json:"page"`
			Preferred string `xml:"Preferred" bson:"preferred" json:"preferred"`
		} `xml:"Term_Source" bson:"termSources" json:"termSources"`
	} `xml:"Term_Sources" bson:"termSource" json:"termSource"`
}

// Source struct
type Source struct {
	SourceID string `xml:"Source_ID" bson:"sourceId" json:"sourceId"`
}

// FindAll finds all subjects.
func (s *Subjects) FindAll(query *mutator.Query) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := database.Connect(collection)

	filter := filter.Subject{}
	filter.Set(query)

	opts := options.Find().SetSkip((query.Page - 1) * query.PageSize).SetLimit(query.PageSize)
	cur, err := c.Find(ctx, filter.Fliter, opts)

	if err != nil {
		return err
	}

	for cur.Next(ctx) {
		subject := Subject{}

		if err := cur.Decode(&subject); err != nil {
			return err
		}

		s.Subjects = append(s.Subjects, subject)
	}

	if err := cur.Err(); err != nil {
		return err
	}

	return cur.Close(ctx)
}

// Find finds a subject.
func (s *Subject) Find(query bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := database.Connect(collection)

	err := c.FindOne(ctx, query).Decode(&s)

	return err
}

// Upsert updates or inserts subjects.
func (s *Subject) Upsert() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	c := database.Connect(collection)

	query := bson.M{"subjectId": s.SubjectID}
	update := bson.M{"$set": s}

	opts := options.Update().SetUpsert(true)
	_, err := c.UpdateOne(ctx, query, update, opts)

	return err
}

// BulkUpsert bulk updates or inserts subjects.
func (s *Subjects) BulkUpsert() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	c := database.Connect(collection)

	models := []mongo.WriteModel{}

	for _, subject := range s.Subjects {
		query := bson.M{"subjectId": subject.SubjectID}
		update := bson.M{"$set": subject}
		model := mongo.NewUpdateOneModel()
		models = append(models, model.SetFilter(query).SetUpdate(update).SetUpsert(true))
	}

	opts := options.BulkWrite().SetOrdered(false)
	_, err := c.BulkWrite(ctx, models, opts)

	return err
}

// PopulateIndex populates the index of subjects.
func (s *Subjects) PopulateIndex() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := database.Connect(collection)

	keys := []string{
		"subjectId",
		"parentRelationship.preferredParents.parentSubjectId",
		"parentRelationship.nonPreferredParents.parentSubjectId",
		"associativeRelationship.associativeRelationships.relatedSubjectId.vpSubjectId",
		"term.preferredTerms.termId",
		"term.nonPreferredTerms.termId",
	}

	models := []mongo.IndexModel{}

	for _, key := range keys {
		model := mongo.IndexModel{
			Keys:    bson.M{key: 1},
			Options: options.Index().SetName(key),
		}
		models = append(models, model)
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err := c.Indexes().CreateMany(ctx, models, opts)

	return err
}
