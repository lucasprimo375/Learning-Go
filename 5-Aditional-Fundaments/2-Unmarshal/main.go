package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  uint8  `json:"age"`
}

func main() {
	dogJSON := `{"name":"Rex","race":"Dalmatian","age":3}`

	var dog1 dog

	fmt.Println(dog1)

	error := json.Unmarshal([]byte(dogJSON), &dog1)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(dog1)

	dog2JSON := `{"name":"Toby","race":"Bulldog"}`

	dog2 := make(map[string]string)

	fmt.Println(dog2)

	error = json.Unmarshal([]byte(dog2JSON), &dog2)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(dog2)
}
