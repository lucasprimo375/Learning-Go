package main

import (
	"fmt"
	"module-example/auxiliar" // module name followed by package name
	"github.com/badoux/checkmail"
)

func main()  {
	fmt.Println("Writing from main file")
	auxiliar.AuxiliarWrite()

	error := checkmail.ValidateFormat("email-test@hotmail.com")
	fmt.Println(error)

	error = checkmail.ValidateFormat("not an e-mail")
	fmt.Println(error)
}