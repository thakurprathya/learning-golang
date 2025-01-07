package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	// pointers are reference to the memory address
	// pointers in go confirms that actual variable is passed on not its copy (passed by reference concept)

	var ptr *int // integer pointer
	fmt.Println("ptr:", ptr)

	mynum := 25
	ptr = &mynum

	fmt.Println("ptr:", ptr)
	fmt.Println("*ptr:", *ptr)

	var ptrToPtr **int = &ptr // pointer to integer pointer

	fmt.Println("ptrToPtr:", ptrToPtr)
	fmt.Println("*ptrToPtr:", *ptrToPtr)
	fmt.Println("**ptrToPtr:", **ptrToPtr)

	fmt.Println("\nUpdating pointer value ( x by 2 )")
	**ptrToPtr = **ptrToPtr * 2
	fmt.Println("**ptrToPtr:", **ptrToPtr)
	fmt.Println("*ptr:", *ptr)
	fmt.Println("mynum:", mynum)
}
