package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")

	// There is no inheritance in golang, no concept of super/parent etc
	me := User{Name: "Prathya", Email: "prathya@mail.dev", Status: false, Age: 21}
	fmt.Println("me:", me)
	fmt.Printf("me: %v\n", me)
	fmt.Printf("me: %+v\n", me) // will print details too
	fmt.Println("me.Name:", me.Name)
}

// first letter capital for public as will be exported and used outside
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
