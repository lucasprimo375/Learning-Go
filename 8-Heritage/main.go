package main

import "fmt"

type person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	person
	course string
	school string
}

func main() {
	fmt.Println("Heritage (of sorts)\n")

	person1 := person{"Lucas", "Primo", 25, 170}
	fmt.Println(person1)

	student1 := student{person1, "Computer Science", "MIT"}
	fmt.Println(student1)
	fmt.Println(student1.name)
	fmt.Println(student1.surname)
	fmt.Println(student1.age)
}
