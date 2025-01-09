package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files")

	write()
	read()
}

func write() {
	content := "This is to be written in a file"

	file, err := os.Create("./firstfile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length of file written:", length)
	file.Close() // can use defer for it just after we create the file and didn't receive any error
}

func read() {
	databytes, err := os.ReadFile("./firstfile.txt")
	checkNilError(err)

	fmt.Println("Text data inside the file:", string(databytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err) // will throw the error & shut down the execution of program
	}
}
