package main

import (
	"fmt"
	"tests/adresses"
)

func main() {
	addressType := adresses.TypeOfAddress("Avenue Paulista")
	fmt.Println(addressType)

	addressType = adresses.TypeOfAddress("Invalid address")
	fmt.Println(addressType)
}
