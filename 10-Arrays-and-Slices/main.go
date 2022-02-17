package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays\n")

	var array [5]int // applies zero value to all items
	fmt.Println(array)

	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	array2 := [5]string{"item1", "item2", "item3", "item4", "item5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // still fixed size
	fmt.Println(array3, "\n")

	fmt.Println("Slices\n")

	slice := []int{}
	fmt.Println(slice)

	slice2 := []int{10, 11, 12, 13}
	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(array3), reflect.TypeOf(slice2)) // arrays and slices are not the same type

	slice2 = append(slice2, 14) // returns new slice with appended item
	fmt.Println(slice2)

	slice2 = array3[1:3] // returns index 1 until 2, i.e. [i:j] returns i, i+1, ..., j-2, j-1 (provided i<j)
	fmt.Println(slice2)

	array3[1] = 500 // pointers in action
	fmt.Println(slice2)
}
