package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the game"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) // os.Stdin specifies reading from keyboard input
	fmt.Println("Enter the input: ")

	// comma ok / error error syntax, go does not have try catch syntax
	input, err := reader.ReadString('\n') // read till end of line, to remove the compiler error (declare, but not used) input, _
	// input, _ := reader.ReadString('\n')
	// if input imp( input, _ ), if error imp( _, err ), if both imp( input, err )
	fmt.Println("input: ", input)
	fmt.Printf("type of input is %T \n", input)
	fmt.Println("error: ", err)

}
