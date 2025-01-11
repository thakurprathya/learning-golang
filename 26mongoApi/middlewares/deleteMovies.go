package middlewares

import (
	"context"
	"fmt"
	"log"
	"mongoApi/database"

	"go.mongodb.org/mongo-driver/bson"
)

// Delete all the movies
func DeleteMovies() {
	res, err := database.Collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted count:", res.DeletedCount)
}
