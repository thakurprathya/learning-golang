package main

import (
	"fmt"
	"log"
	"mongoApi/router"
	"net/http"
)

func main() {
	fmt.Println("mongoApi in golang")

	// for mongodb connection -> go get go.mongodb.org/mongo-driver/mongo  :driver for go
	// for handling env file -> go get github.com/joho/godotenv :https://github.com/joho/godotenv

	fmt.Println("Server is starting...")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
