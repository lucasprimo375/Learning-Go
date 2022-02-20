package main

import (
	"bytes"
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
	dog1 := dog{"Rex", "Dalmatian", 3}

	fmt.Println(dog1)

	dog1Json, error := json.Marshal(dog1)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(dog1Json)

	fmt.Println(bytes.NewBuffer(dog1Json))

	dog2 := map[string]string{
		"name": "Toby",
		"race": "Bulldog",
	}

	dog2Json, error := json.Marshal(dog2)

	fmt.Println(dog2Json)

	fmt.Println(bytes.NewBuffer(dog2Json))
}
