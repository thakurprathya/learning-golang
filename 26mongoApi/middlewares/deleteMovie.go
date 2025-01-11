package middlewares

import (
	"context"
	"fmt"
	"log"
	"mongoApi/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete movie based on id
func DeleteMovie(movieId string) {
	_id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	// https://stackoverflow.com/questions/64281675/bson-d-vs-bson-m-for-find-queries
	filter := bson.M{"_id": _id}
	res, err := database.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted count:", res.DeletedCount)
}
