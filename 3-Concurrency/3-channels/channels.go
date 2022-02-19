package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Hello World!", channel)

	fmt.Println("After calling function write")

	/*for {
		message, channelIsOpen := <-channel // channel waits until it gets a value and then put it into the variable

		if !channelIsOpen {
			break
		}

		fmt.Println(message)
	}*/

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("End of execution")
}

func write(text string, channel chan string) {
	time.Sleep(time.Second * 5)

	for i := 0; i < 5; i++ {
		channel <- text // sending a value inside the channel
		time.Sleep(time.Second)
	}

	close(channel)
}
