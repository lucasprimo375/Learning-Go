package main

import (
	"fmt"
)

func closure() func() {
	text := "Inside closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	fmt.Println("Closure")

	text := "Inside main function"
	fmt.Println(text)

	newFunction := closure()
	newFunction()
}
