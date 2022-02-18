package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Interface as Generics")

	generic("string")
	generic(10)
	generic(true)

	fmt.Println(1, 2, "text", true, false, float64(123.456))

	mymap := map[interface{}]interface{}{
		1:    "string",
		true: float64(1.25),
	}
	fmt.Println(mymap)
}
