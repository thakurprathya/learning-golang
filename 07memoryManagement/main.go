package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go Memory Management")

	// Go automatically mangages its memory
	// there are 2 keywords use for assiging memory
	// 1. new - assign memory but does not initialize variable, zeroed storage(can not put any data initially), will get memory address
	// 2. make - assign memory and initialize variable, non-zeroed storage, will get memory address

	// Go offers automatic garbage collection
	// anything in or out of scope is added or removed accordingly
	// runtime library

	currentCPUs := runtime.NumCPU()
	fmt.Println("current number of cpu's :", currentCPUs)
}
