package main

import "fmt"

const LoginToken string = "afdasfdasf" // const variable, val cannot be changed, imp thing is as it starts with capital letter it is equivalent to public: const loginToken from any other language

func main() {
	var username string = "First String"
	fmt.Println(username)
	fmt.Printf("Variable username = %s, is of type %T \n\n", username, username) // %s for string, %T for type of variable

	var check bool = false
	fmt.Println(check)
	fmt.Printf("Variable check = %t, is of type %T \n\n", check, check) // %t for bool, %T for type of variable

	var smallUnsignedValue uint8 = 255 // 0 to 255, another alias byte
	fmt.Println(smallUnsignedValue)
	fmt.Printf("Variable smallUnsignedValue = %d, is of type %T \n\n", smallUnsignedValue, smallUnsignedValue) // %d for integers, %T for type of variable

	var smallSignedValue int8 = -128 // -128 to 127
	fmt.Println(smallSignedValue)
	fmt.Printf("Variable smallSignedValue = %d, is of type %T \n\n", smallSignedValue, smallSignedValue) // %d for integers, %T for type of variable

	var integer int = 1234123414 // int32 only, another alias for int32 is rune
	fmt.Println(integer)
	fmt.Printf("Variable integer = %d, is of type %T \n\n", integer, integer) // %d for integers, %T for type of variable

	var smallFloat float32 = 234.214312341 // store only upto 5 decimal places
	fmt.Println(smallFloat)
	fmt.Printf("Variable smallFloat = %f, is of type %T \n\n", smallFloat, smallFloat) // %f for decimal, %T for type of variable

	var preciseFloat float64 = 234.214312341 // store only upto 5 decimal places
	fmt.Println(preciseFloat)
	fmt.Printf("Variable preciseFloat = %.10f, is of type %T \n\n", preciseFloat, preciseFloat) // %f for decimal, %T for type of variable

	// Default Values
	var test int
	fmt.Println(test)
	fmt.Println()

	// implicit types
	var str = "my string"
	fmt.Println(str)
	fmt.Println()
	// str = 5  // this is not allowed

	// alternative for var (only valid inside a function/method) : walrus operator
	variable := 513.2
	fmt.Println(variable)
	fmt.Println()

	// const variable
	fmt.Println(LoginToken)
}
