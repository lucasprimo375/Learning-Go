package main

import "fmt"

func main() {
	channel := make(chan string, 2) // second argument is channel capacity. channel only blocks after capacity is reached

	channel <- "Hello World!"
	channel <- "Programming with go"
	//channel <- "Hello there!" // deadlock

	message := <-channel
	message2 := <-channel
	//message3 := <-channel // deadlock

	fmt.Println(message)
	fmt.Println(message2)
	//fmt.Println(message3)
}
