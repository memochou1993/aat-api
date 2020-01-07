package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri        = "mongodb://localhost:27017"
	database   = "thesaurus"
	collection = "subjects"
)

var err error
var client *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(uri)
	if client, err = mongo.Connect(ctx, opts); err != nil {
		log.Fatalln(err.Error())
	}
}

// Upsert updates or inserts a resource.
func Upsert(query interface{}, update interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := client.Database(database).Collection(collection)

	opts := options.Update().SetUpsert(true)
	_, err := c.UpdateOne(ctx, query, update, opts)

	return err
}
