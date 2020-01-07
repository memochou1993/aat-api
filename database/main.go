package database

import (
	"log"

	"github.com/globalsign/mgo"
)

const (
	host       = "localhost:27017"
	source     = ""
	username   = ""
	password   = ""
	db         = "thesaurus"
	collection = "subjects"
)

var session *mgo.Session

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Source:   source,
		Username: username,
		Password: password,
	}

	s, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Fatalln(err.Error())
	}

	session = s
}

func connect() (*mgo.Session, *mgo.Collection) {
	s := session.Copy()
	c := s.DB(db).C(collection)

	return s, c
}

// Upsert updates or inserts one or more documents.
func Upsert(selector interface{}, update interface{}) error {
	s, c := connect()
	defer s.Close()

	_, err := c.Upsert(selector, update)

	return err
}
