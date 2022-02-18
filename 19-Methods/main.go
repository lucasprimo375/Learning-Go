package main

import (
	"fmt"
)

func write() {
	fmt.Println("Writing...")
}

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Println("Inside save method")

	fmt.Printf("Saving user %s data into database\n", u.name)
}

func (u user) isAgeOfMajority() bool {
	return u.age >= 18
}

func (u *user) getOlder() {
	u.age++
}

func main() {
	fmt.Println("Methods")

	write()

	user1 := user{"Lucas", 25}
	fmt.Println(user1)

	user1.save()

	user2 := user{"Robert", 75}
	user2.save()

	fmt.Println(user1.isAgeOfMajority())
	fmt.Println(user2.isAgeOfMajority())

	fmt.Println(user1.age)
	user1.getOlder()
	fmt.Println(user1.age)
}
