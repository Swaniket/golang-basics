package controller

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connString = "mongodb://localhost:27017"
const dbName = "GoTestDB"
const collectionName = "movies"

// Get the ref of mongo collection
var collection *mongo.Collection

// Connect with Mongo
func init() { // Init is an initialization method in go which only runs when the program is initialized
	// Client options
	clientOptions := options.Client().ApplyURI(connString)

	// Connect to mongo
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Mongo Connection Success!")

	collection = client.Database(dbName).Collection(collectionName)

	// Collection Instance
	fmt.Println("Collection ref is ready")
}
