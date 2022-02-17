package main

import "fmt"

// function declation
func sum(number1 int, number2 int) int {
	return number1 + number2
}

// multiple returns
func calculate(number1, number2 int) (int, int) { // single declaration
	sumresult := number1 + number2
	subtraction := number1 - number2

	return sumresult, subtraction
}

func main() {
	// caling function
	sumresult := sum(1, 2)
	fmt.Println(sumresult)

	// function is a type too and can be attributed to a variable
	var f = func(text string) string {
		fmt.Println(text)
		return text
	}

	result := f("text example")
	fmt.Println(result)

	// calling function with multiple returns
	sumresult, subtraction := calculate(1, 2)
	fmt.Println(sumresult, subtraction)

	sumresult2, _ := calculate(1, 2) // if we have no need for the second return. _ will ignore it
	fmt.Println(sumresult2)

	_, subtraction2 := calculate(1, 2) 
	fmt.Println(subtraction2)
}