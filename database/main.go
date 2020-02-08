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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(uri)
	if client, err = mongo.Connect(ctx, opts); err != nil {
		log.Fatalln(err.Error())
	}
}

// Connect connects to the database.
func Connect(collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}
