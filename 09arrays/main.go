package main

import "fmt"

func main() {
	fmt.Println("Array in Golang")

	var firstArr [3]int = [3]int{1, 2, 3}
	fmt.Println("firstArr:", firstArr)

	// array operations
	firstArr[0] = 100
	fmt.Println("firstArr:", firstArr)
	fmt.Println("len(firstArr):", len(firstArr))
	fmt.Printf("type of firstArr: %T", firstArr)
}
