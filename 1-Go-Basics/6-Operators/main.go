package main

import "fmt"

func main() {
	// arithmetic operators: +, -, *, /, % (mod)
	fmt.Println("arithmetic operators")
	sum := 1 + 2 // int16 can not be sumed with int32. The same goes for every different types, even if both are numbers
	subtraction := 1 - 2
	multiplication := 1 * 2
	division := 1 / 2
	mod := 1 % 2
	fmt.Println(sum, subtraction, multiplication, division, mod, "\n")

	// attribution
	fmt.Println("attribution operators")
	var string1 string = "string1"
	string2 := "string2"
	fmt.Println(string1, string2, "\n")

	// relational
	fmt.Println("relational operators")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2, "\n")

	// logical
	fmt.Println("logical operators")
	fmt.Println(true && true) // && also works
	fmt.Println(true || true) // || also works
	fmt.Println(!true, "\n")  // ! also works

	// unary operatos
	fmt.Println("unary operators")
	number := 10
	number++
	fmt.Println(number)

	number += 10
	fmt.Println(number)

	number-- // --number and ++number not possible
	fmt.Println(number)

	number -= 10
	fmt.Println(number)

	number *= 10
	fmt.Println(number)

	number /= 5
	fmt.Println(number)

	number %= 3
	fmt.Println(number, "\n")

	// <x> ? <y> : <x> does work
}
