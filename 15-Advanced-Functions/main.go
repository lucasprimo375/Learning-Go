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

func fibonacci(index uint) uint {
	if index <= 1 {
		return index
	}

	return fibonacci(index-1) + fibonacci(index-2)
}

func func1() {
	fmt.Println("Executing func1")
}

func func2() {
	fmt.Println("Executing func2")
}

func isApproved(grade1, grade2 float64) bool {
	defer fmt.Println("Final grande calculated. Result will be returned")

	fmt.Println("Entering function to check if student is approved")

	finalGrade := (grade1 + grade2) / 2.0

	if finalGrade >= 6 {
		return true
	}

	return false
}

func recoverExecution() {
	fmt.Println("Trying to recover execution")

	if r := recover(); r != nil { // recover returns nil if execution is not in panic
		fmt.Println("Execution recovered successfully")
	}
}

func isApproved2(grade1, grade2 float64) bool {
	defer recoverExecution()

	finalGrade := (grade1 + grade2) / 2.0

	if finalGrade > 6 {
		return true
	} else if finalGrade < 6 {
		return false
	}

	panic("Final grande is exactly 6!")
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

	fmt.Println("Recursive Functions")
	var fibonacciResult uint = fibonacci(10)
	fmt.Println(fibonacciResult)

	fmt.Println("Defer clause")
	defer func1() // postpone the execution of func1 until possible
	func2()

	fmt.Println(isApproved(7, 8))

	// Panic and Recover
	fmt.Println("isApproved2=", isApproved2(6, 6))
	fmt.Println("Post execution")
}
