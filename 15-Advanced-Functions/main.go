package main

import "fmt"

// named return
func calculate(number1, number2 int) (sum int, subtraction int) {
	sum = number1 + number2
	subtraction = number1 - number2
	return
}

// variadic function
func sumFunc(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func write(text string, numbers ...int) { // can have only one variadic parameter and it has to be the last one
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	fmt.Println("Advanced Functions")

	sum, subtraction := calculate(1, 2)
	fmt.Println(sum, subtraction)

	sumTotal := sumFunc(1, 2, 3, 4, 5, 6, 7, 8) // can also give zero arguments
	fmt.Println(sumTotal)

	write("string", 1, 2, 3, 4, 5, 6, 7, 8)

	// anonymous functions
	text := "sample text"
	text1 := func(text string) string {
		fmt.Println("Anonymous function:", text)
		return fmt.Sprintf("Got %s", text)
	}(text)
	fmt.Println("Anonymous function returns:", text1)
}
