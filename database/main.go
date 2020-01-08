package database

import (
	"context"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload" // initialize
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	uri      = os.Getenv("DB_URI")
	database = os.Getenv("DB_DATABASE")
	err      error
	client   *mongo.Client
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(uri)
	if client, err = mongo.Connect(ctx, opts); err != nil {
		log.Fatalln(err.Error())
	}
}

// Upsert updates or inserts a document.
func Upsert(collection string, query interface{}, update interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := client.Database(database).Collection(collection)

	opts := options.Update().SetUpsert(true)
	_, err := c.UpdateOne(ctx, query, update, opts)

	return err
}

// BulkUpsert bulk updates or inserts documents.
func BulkUpsert(collection string, models []mongo.WriteModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	c := client.Database(database).Collection(collection)

	opts := options.BulkWrite().SetOrdered(false)
	_, err := c.BulkWrite(ctx, models, opts)

	return err
}
