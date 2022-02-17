package main

import "fmt"

func main() {
	fmt.Println("Maps\n")

	user := map[string]string{
		"name":    "Lucas",
		"surname": "Primo",
	}

	fmt.Println("Map user: ", user)
	fmt.Println("Value of key name in map user: ", user["name"], "\n")

	fmt.Println("Nested maps\n")

	user2 := map[string]map[string]string{
		"name": {
			"first":      "Robert",
			"middlename": "Kardashian",
		},
		"course": {
			"school": "MIT",
			"name":   "Computer Science",
		},
	}

	fmt.Println("map user2: ", user2, "\n")

	fmt.Println("Deleting key 'course' from map...\n")
	delete(user2, "course")
	fmt.Println("map user2 after deletion: ", user2, "\n")

	fmt.Println("Adding new key, value pair into from user2...")
	user2["sign"] = map[string]string{
		"name": "Twins",
	}
	fmt.Println("user2: ", user2, "\n")
}
