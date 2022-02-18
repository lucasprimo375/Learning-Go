package main

import "fmt"

func main() {
	fmt.Println("Pointers\n")

	var var1 int = 10
	var var2 int = var1
	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	var2++
	fmt.Println(var1, var2)

	var var3 int = 100
	var pointer *int
	fmt.Println(var3, pointer)

	pointer = &var3
	fmt.Println(var3, pointer)

	fmt.Println(var3, *pointer) // dereferencing

	var3 = 150
	fmt.Println(var3, pointer)
	fmt.Println(var3, *pointer)
}
