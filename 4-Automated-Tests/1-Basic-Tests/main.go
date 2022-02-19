package main

import (
	"fmt"
	"tests/addresses"
)

func main() {
	addressType := addresses.TypeOfAddress("Avenue Paulista")
	fmt.Println(addressType)

	addressType = addresses.TypeOfAddress("Invalid address")
	fmt.Println(addressType)
}
