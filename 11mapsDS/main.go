package main

import "fmt"

func main() {
	fmt.Println("Maps in golang (HashMap / Hashtable)")

	keypair := make(map[string]string) // declaring a map of key value string type
	keypair["f"] = "first"
	keypair["s"] = "second"
	keypair["t"] = "third"

	fmt.Println("keypair:", keypair)
	fmt.Println("keypair[f]:", keypair["f"])

	// deleting a key-value pair
	delete(keypair, "f")
	fmt.Println("keypair:", keypair)

	// for loop in maps (similar to for each loop)
	for key, value := range keypair {
		// fmt.Printf("keypair[%s]: %s\n", key, value)
		fmt.Printf("keypair[%v]: %v\n", key, value)
	}
	// again here if anything not imp replace with _, value | key, _
}
