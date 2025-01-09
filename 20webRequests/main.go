package main

import (
	"fmt"
	"io"
	"net/http"
)

const myurl = "https://reqres.in/api/users"

func main() {
	fmt.Println("webRequests in golang")

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // necessary in go to close the connection

	fmt.Printf("Type of response is %T\n", response)

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
