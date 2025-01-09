package main

import "fmt"

func main() {
	fmt.Println("methods")

	// There is no inheritance in golang, no concept of super/parent etc
	me := User{Name: "Prathya", Email: "prathya@mail.dev", Status: false, Age: 21, online: true}
	// fmt.Println("me:", me)
	// fmt.Printf("me: %v\n", me)
	fmt.Printf("me: %+v\n", me) // will print details too

	me.GetStatus() // calling the method

	me.SetEmail("thakur@mail.dev") // call by value kind
	fmt.Println("me.Email:", me.Email)

	me.SetEmail2("thakur@mail.dev") // call by reference kind
	fmt.Println("me.Email:", me.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	online bool // not exportable starting with lowercase
}

// First Method
func (u User) GetStatus() { // starting with uppercase to make exportable
	fmt.Println("Is User Active:", u.Status)
}

func (u User) SetEmail(e string) { // recieves a copy
	fmt.Println("current email:", u.Email)
	u.Email = e
	fmt.Println("updated email:", u.Email)
}

func (u *User) SetEmail2(e string) { //  recieves actual object
	fmt.Println("current email:", u.Email)
	u.Email = e
	fmt.Println("updated email:", u.Email)
}
