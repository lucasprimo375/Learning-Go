package main

import (
	"fmt"
	"module-example/auxiliar" // module name followed by package name
)

func main()  {
	fmt.Println("Writing from main file")
	auxiliar.AuxiliarWrite()
}