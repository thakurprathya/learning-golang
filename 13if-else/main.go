package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("if-else Control Flow in Go")

	fmt.Println("Enter your age: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	age, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	// strict syntax cannot take { or else to next lines
	if age >= 21 {
		fmt.Println("Legal Drinking Age")
	} else if age >= 18 {
		fmt.Println("You can try!")
	} else {
		fmt.Println("No way, Get Out!!")
	}

	// variable assigning and checking (used when handling requests)
	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else if num > 10 {
		fmt.Println("num is greater than 10")
	} else {
		fmt.Println("num is 10")
	}
}
