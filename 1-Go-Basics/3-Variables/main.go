package main

import "fmt"

func main() {
	// two options to declare variables
	
	var variable1 string = "Variable 1" // explicit type
	variable2 := "Variable 2" // implicit type
	
	fmt.Println(variable1)
	fmt.Println(variable2)

	// declaring multiple variables

	var (
		variable3 string = "Variable 3"
		variable4 string = "Variable 4"
	)

	fmt.Println(variable3, variable4)

	variable5, variable6 := "Variable 5", "Variable 6"

	fmt.Println(variable5, variable6)

	// constants

	const constant1 string = "Constant 1"
	fmt.Println(constant1)

	// exchanging variables values

	variable5, variable6 = variable6, variable5
	fmt.Println(variable5, variable6)
}