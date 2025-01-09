package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://reqres.in:3000/api/users?page=2&data=something+nothing"

func main() {
	fmt.Println("urls in golang")

	// parsing the URL
	result, _ := url.Parse(myurl)

	fmt.Println("result.Scheme:", result.Scheme)
	fmt.Println("result.Host:", result.Host)
	fmt.Println("result.Path:", result.Path)
	fmt.Println("result.Port():", result.Port())
	fmt.Println("result.RawQuery:", result.RawQuery)

	qParams := result.Query() // returns a map
	fmt.Println("qParams:", qParams)
	for key, val := range qParams {
		fmt.Printf("qParams[%v] = %v\n", key, val)
	}

	// constructing a URL
	query := url.Values{}
	query.Add("id", "asfda")

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "sample.com",
		Path:     "/api/docs",
		RawQuery: query.Encode(),
	}

	newUrl := partsOfUrl.String()
	fmt.Println("newUrl:", newUrl)
}
