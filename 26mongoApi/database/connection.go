package database

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload" // automatically loads root env file
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseURL = os.Getenv("MONGO_URI")
var databaseName = "movieflix"
var collectionName = "watchlist"

var Collection *mongo.Collection

// init is a special function, runs only for the first time when initializing the application
func init() {
	// Create MongoDB client options by configuring the connection URI
	clientOption := options.Client().ApplyURI(databaseURL)

	// Establish a connection to MongoDB using the configured options, for understanding context :https://pkg.go.dev/context
	client, err := mongo.Connect(context.TODO(), clientOption) // as unclear what context to use, we use TODO context
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to Mongodb success!!")

	// reference to collection
	Collection = client.Database(databaseName).Collection(collectionName)
	fmt.Println("collection instance ready!")
}
