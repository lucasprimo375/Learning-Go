package main

import (
	"fmt"
)

func invertSign(number int) int {
	return -number
}

func invertSignWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	fmt.Println("Functions with Pointers")

	number := 10
	invertedNumber := invertSign(number)
	fmt.Println("invertedNumber =", invertedNumber)
	fmt.Println("number =", number)

	number1 := 40
	fmt.Println("number2 =", number1)
	invertSignWithPointer(&number1)
	fmt.Println("number2 =", number1)
}
