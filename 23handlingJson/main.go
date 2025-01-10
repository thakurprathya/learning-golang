package main

import (
	"encoding/json"
	"fmt"
)

type course struct { // startswith lowercase, non exportable
	Name     string `json:"coursename"` // this flag will rename field to coursename for json
	Price    int
	Platform string   `json:"site"`
	Password string   `json:"-"`              // this flag means don't consume this field in json
	Tags     []string `json:"tags,omitempty"` // this flag means rename to tags & don't consider empty one
}

func main() {
	fmt.Println("handlingJson in golang")

	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"MERN Stack", 5999, "Udemy", "pass1234", []string{"web", "frontend", "backend"}},
		{"React Js", 2999, "Udemy", "pass2345", []string{"web", "frontend"}},
		{"Node Js", 2999, "Udemy", "pass3456", []string{"web", "backend"}},
		{"Flask", 2999, "Udemy", "passss", nil},
	}

	// converting to json
	// method1
	jsonData1, err := json.Marshal(courses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData1) // if directly print will return byte of numbers

	//method2
	jsonData2, err := json.MarshalIndent(courses, "", "\t") // data will be indented based on \t, with nothing as prefix
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData2) // if directly print will return byte of numbers
}

func DecodeJson() {
	// format of request data is slice of byte
	sampleJson := []byte(`
		{
			"coursename": "React Js",
			"Price": 2999,
			"site": "Udemy",
			"tags": ["web","frontend"]
        }
	`)

	// validating json & populating as per interfaces (course struct)
	checkValid := json.Valid(sampleJson)
	var responseCourse course

	if checkValid {
		fmt.Println("json is valid")
		json.Unmarshal(sampleJson, &responseCourse) // passing by reference as method
		fmt.Printf("%#v\n", responseCourse)         // special flag for printing interfaces
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// general case where decode data to key value pairs
	// below is the general datatype used for storing request type data
	var data map[string]interface{} // as value is not guarante marking it as interface so everything valid (bool, int, string, array)
	json.Unmarshal(sampleJson, &data)
	fmt.Printf("%#v\n", data) // special flag for printing interfaces

	for k, v := range data {
		fmt.Printf("key[%v]: %v => type(%T)\n", k, v, v)
	}
}
