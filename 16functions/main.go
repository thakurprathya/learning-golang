package main

import "fmt"

func add(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Println("functions in go")

	greet()
	fmt.Println("sum of 123 and 251:", add(123, 251))

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("sum of s:", adder(s...))

	n1 := 123
	n2 := 513
	n3 := 50
	fmt.Println("sum of n1, n2 & n3:", adder(n1, n2, n3))

	// multiple return function
	n, str := check()
	fmt.Println("n:", n)
	fmt.Println("str:", str)
}

func greet() {
	fmt.Println("Morning")
}

func adder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func check() (int, string) {
	return 5, "hello world"
}
