package main

import "fmt"

func main() {
	fmt.Println("Control Structures")

	number := 10

	if number > 15 { // even if it's just one command inside of 'if', {} must be present
		fmt.Println("More than 15")
	} else {
		fmt.Println("Less or equal than 15")
	}

	if number2 := number; number2 > 0 { // number2 is inside of the 'if-else' scope
		fmt.Println("number2 is more than 0")
	} else if number2 == 0 {
		fmt.Println("number2 is equal to 0")
	} else {
		fmt.Println("number2 is less than 0")
	}
}
