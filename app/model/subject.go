package model

import (
	"github.com/memochou1993/thesaurus/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collection = "subjects"
)

// Vocabulary struct
type Vocabulary struct {
	Title    string    `xml:"Title,attr"`
	Date     string    `xml:"Date,attr"`
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
	ParentSubjectID  string `xml:"Parent_Subject_ID" bson:"parentSubjectID" json:"parentSubjectID"`
	RelationshipType string `xml:"Relationship_Type" bson:"relationshipType" json:"relationshipType"`
	HistoricFlag     string `xml:"Historic_Flag" bson:"historicFlag" json:"historicFlag"`
	ParentString     string `xml:"Parent_String" bson:"parentString" json:"parentString"`
	HierRelType      string `xml:"Hier_Rel_Type" bson:"hierRelType" json:"hierRelType"`
}

// AssociativeRelationship struct
type AssociativeRelationship struct {
	RelationshipType string `xml:"Relationship_Type" bson:"relationshipType" json:"relationshipType"`
	RelatedSubjectID struct {
		VPSubjectID string `xml:"VP_Subject_ID" bson:"vpSubjectID" json:"vpSubjectID"`
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

// Upsert updates or inserts a subject.
func (m *Subject) Upsert(query interface{}, subject interface{}) error {
	return database.Upsert(collection, query, subject)
}

// BulkUpsert bulk updates or inserts subjects.
func (m *Subject) BulkUpsert(subjects []Subject) error {
	models := []mongo.WriteModel{}

	for _, subject := range subjects {
		query := bson.M{"subjectId": subject.SubjectID}
		update := bson.M{"$set": subject}
		model := mongo.NewUpdateOneModel()
		models = append(models, model.SetFilter(query).SetUpdate(update).SetUpsert(true))
	}

	return database.BulkUpsert(collection, models)
}
