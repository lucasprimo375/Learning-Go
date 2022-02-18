package main

import "fmt"

type address struct {
	street  string
	zipcode uint8
}

type user struct { // zero value of struct sets zero value for each field
	name        string
	age         uint8
	useraddress address
}

func main() {
	fmt.Println("Structs", "\n")

	var u user
	u.name = "Lucas"
	u.age = 25
	fmt.Println(u)

	address1 := address{"Streetname", 123}

	u2 := user{"Robert", 71, address1}
	fmt.Println(u2)

	u3 := user{name: "Alan"} // in case we do not have all data to put inside the struct
	fmt.Println(u3)

	u4 := user{age: 50}
	fmt.Println(u4)
}
