package main

import "fmt"

/*
defer is like a "do it later" command - it schedules a function to run
right before the current function returns. Think of it like a reminder
note that says "don't forget to do this before you leave!"

Some common uses of defer:
- Closing files after you're done with them
- Cleaning up resources
- Making sure something runs even if there's an error

Important notes:
- Deferred functions run in LIFO order (Last In, First Out)
- Like a stack of plates - the last defer you add is the first one that runs
- The defer statement is executed immediately, but the function call is delayed

Example:
    defer fmt.Println("bye")   // This prints last
    defer fmt.Println("hi")    // This prints second
    fmt.Println("hello")       // This prints first
*/

func main() {
	fmt.Println("defers")

	defer fmt.Println("third line") // this line simply means move this line just above end of function (just above } )
	defer fmt.Println("second line")
	fmt.Println("first line")

	// inside defer use LIFO last in first out
	reverse()
}

func reverse() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
