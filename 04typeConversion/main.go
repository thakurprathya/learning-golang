package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a Integer: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Input: ", input)

	// String Manipulation
	updatedInput := strings.TrimSpace(input)

	// Type conversion
	numInput, err := strconv.ParseFloat(updatedInput, 64) // variable to parse and bitSize

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Numeric Input: ", numInput)
	}
}
