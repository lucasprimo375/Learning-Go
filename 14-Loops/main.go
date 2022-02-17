package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		fmt.Println(i)
		//time.Sleep(time.Second)
		i++
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		//time.Sleep(time.Second)
	}

	for j := 0; j < 10; j += 2 {
		fmt.Println(j)
		//time.Sleep(time.Second)
	}

	names := [3]string{"Lucas", "Robert", "Alan"}

	for i, name := range names { // we can use _ if we have no interest in either of the values
		fmt.Println(i, ": ", name)
	}

	names2 := []string{"Lucas", "Robert", "Alan"}

	for i, name := range names2 {
		fmt.Println(i, ": ", name)
	}

	for i, letter := range "STRING" {
		fmt.Println(i, ": ", letter) // printing a char, which is the ASCII code
	}

	for i, letter := range "STRING" {
		fmt.Println(i, ": ", string(letter))
	}

	users := map[string]string{
		"name":    "Lucas",
		"surname": "Primo",
	}

	for key, value := range users {
		fmt.Println(key, " -> ", value)
	}

	// infinite loop
	// for {

	// }
}
