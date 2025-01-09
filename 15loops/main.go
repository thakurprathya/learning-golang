package main

import "fmt"

func main() {
	fmt.Println("loops in go")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("days:", days)

	// base for loop
	for i := 0; i < len(days); i++ { // no ++i exist here 😞
		fmt.Println(days[i])
	}

	// for range loop
	for i := range days {
		fmt.Println(days[i])
	}

	// for each type loop
	for _, day := range days { // not needed index
		fmt.Println(day)
	}

	// for loop as while loop
	j := 0
	for j < len(days) {
		fmt.Println(days[j])
		j++
	}

	// break, continue & goto statements similar to c/cpp

	//goto
	i := 1
	for ; i < 10; i++ {
		fmt.Println(i)
		if i == 7 {
			goto point
		}
	}

point:
	fmt.Println("this is checkpoint")
}
