package main

import (
	"fmt"
)

var n int

// executed before main
func init() {
	fmt.Println("Executing init function")
	n = 10
}

func main() {
	fmt.Println("Executing main function")
	fmt.Println(n)
}
