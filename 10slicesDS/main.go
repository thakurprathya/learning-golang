package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in go")

	// slices are under the hood arrays, abstracted
	var firstSlice []int = []int{1, 22, 3}
	fmt.Println("firstSlice:", firstSlice)
	fmt.Println("len(firstSlice):", len(firstSlice))
	fmt.Printf("type of firstSlice: %T\n", firstSlice)

	// cannot perform array like operations, index modification
	firstSlice = append(firstSlice, 42, 5, 16)
	fmt.Println("firstSlice:", firstSlice)
	fmt.Println("len(firstSlice):", len(firstSlice))

	firstSlice = firstSlice[2:]
	fmt.Println("firstSlice:", firstSlice)

	firstSlice = append(firstSlice[2:], 14, 8)
	fmt.Println("firstSlice:", firstSlice)

	// slice using make
	newSlice := make([]string, 2)
	newSlice[0] = "first"
	newSlice[1] = "second"
	// newSlice[2] = "third" // value over defined size will raise out of bound error
	newSlice = append(newSlice, "third", "fourth")
	fmt.Println("newSlice:", newSlice)
	fmt.Printf("type of newSlice: %T\n", newSlice)

	// sorting
	fmt.Println("before sort firstSlice:", firstSlice)
	sort.Ints(firstSlice)
	fmt.Println("after sort firstSlice:", firstSlice)
	fmt.Println("is firstSlice sorted:", sort.IntsAreSorted(firstSlice))
}
