package main

import (
	"fmt"
	"errors"
)

func main() {
	// types of integer by size: int8, int16, int32, int64

	var number int16 = 100;
	fmt.Println(number)

	// if not specified, Go will base the size of the machine architecture
	var number1 int = 1;
	fmt.Println(number1)

	// unsigned int. also has size for 8, 16, 32, 64
	var unsignedint uint = 10;
	fmt.Println(unsignedint)

	// alias
	var number3 rune = 32 // rune is alias for int32
	var number4 byte = 8 // byte is alias for int 8
	fmt.Println(number3, number4)

	// floating point
	var number5 float32 = 1.23
	var number6 float64 = 11.2322324325
	fmt.Println(number5, number6)

	number7 := 1.23 // implicit float declaration gets size for architecture of the machine
	fmt.Println(number7)

	// strings
	var string1 string = "String 1"
	fmt.Println(string1)

	string2 := "String 2"
	fmt.Println(string2)

	//char
	char := 'B'
	fmt.Println(char) // it prints the correspondent ASCII code

	// zero value: value given to unitialized variable
	var text string
	fmt.Println(text) // empty string

	var intnumber int16
	fmt.Println(intnumber) // number zero for any type of number

	// boolean
	var boolean1 bool = true // or false. also, zero value is false
	fmt.Println(boolean1)

	// error
	var error1 error
	fmt.Println(error1) // zero value is nil

	var error2 error = errors.New("Internal Error")
	fmt.Println(error2)
}