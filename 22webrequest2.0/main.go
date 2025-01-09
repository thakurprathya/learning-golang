package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const baseurl = "http://localhost:3000"

func main() {
	fmt.Println("webrequest2.0 in golang")

	GetRequest(baseurl + "/get")
	PostRequest(baseurl + "/post")
	PostRequest2(baseurl + "/form")
}

func GetRequest(myurl string) {
	response, err := http.Get(myurl)
	checkNilError(err)

	defer response.Body.Close() // necessary in go to close the connection

	statusCode := response.StatusCode
	contentLength := response.ContentLength
	fmt.Println("statusCode:", statusCode)
	fmt.Println("contentLength:", contentLength)

	// method 1: for handling bytes
	// databytes, err := io.ReadAll(response.Body)
	// checkNilError(err)
	// content := string(databytes)
	// fmt.Println(content)

	// method 2: for handling bytes
	var responseString strings.Builder
	databytes, err := io.ReadAll(response.Body)
	checkNilError(err)
	byteCount, _ := responseString.Write(databytes)
	fmt.Println("byteCount:", byteCount)
	fmt.Println(responseString.String())
}

func PostRequest(myurl string) {
	// fake json payload
	requestBody := strings.NewReader(`
        {
            "username" : "admin",
            "id" : "#23",
            "email" : "admin@mail.com"
        }
    `)

	response, err := http.Post(myurl, "application/json", requestBody)
	checkNilError(err)
	defer response.Body.Close() // necessary in go to close the connection

	databytes, err := io.ReadAll(response.Body)
	checkNilError(err)
	content := string(databytes)
	fmt.Println(content)
}

func PostRequest2(myurl string) {
	// form data
	data := url.Values{}
	data.Add("username", "tester")
	data.Add("id", "#11")
	data.Add("email", "tester@mail.com")

	response, err := http.PostForm(myurl, data)
	checkNilError(err)
	defer response.Body.Close() // necessary in go to close the connection

	databytes, err := io.ReadAll(response.Body)
	checkNilError(err)
	content := string(databytes)
	fmt.Println(content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
