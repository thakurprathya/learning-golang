package main

import (
	"fmt"
	"math/rand"
	"time"
	// "crypto/rand"
)

func main() {
	fmt.Println("switch case in Golang")

	// math random
	// Initializes the random number generator using current time as seed to ensure different random numbers on each program execution
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // as time is not constant we use it as source for random algorithm
	diceNumber := r.Intn(6) + 1                          // will generate a number between 0 to 5 and adding 1

	// crypto random, more random more accurate
	// diceNumber, err := rand.Int(rand.Reader, big.NewInt(6)) // will generate a number between 0 to 5

	switch diceNumber {
	case 1:
		fmt.Println("dice rolls 1")
	case 2:
		fmt.Println("dice rolls 2")
	case 3:
		fmt.Println("dice rolls 3")
		fallthrough // if recieved 3 will move to next case too (fall to next case)
	case 4:
		fmt.Println("dice rolls 4")
		fallthrough
	case 5:
		fmt.Println("dice rolls 5")
	case 6:
		fmt.Println("dice rolls 6")
	default:
		fmt.Println("will never reach this case")
	}

}
