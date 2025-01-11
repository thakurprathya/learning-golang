package middlewares

import (
	"context"
	"fmt"
	"log"
	"mongoApi/database"
	"mongoApi/models"
)

func AddMovie(movie models.WatchMovie) interface{} {
	success, err := database.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted Movie in db with id:", success.InsertedID)
	return success.InsertedID
}
