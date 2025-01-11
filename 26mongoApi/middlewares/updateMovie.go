package middlewares

import (
	"context"
	"fmt"
	"log"
	"mongoApi/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateMovie(movieId string) {
	_id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	// https://stackoverflow.com/questions/64281675/bson-d-vs-bson-m-for-find-queries
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := database.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", res.ModifiedCount)
}
