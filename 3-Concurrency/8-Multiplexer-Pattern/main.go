package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(write("Hello world!"), write("Programming with go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}

}

func multiplexer(inputChannel1, inputChannel2 <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-inputChannel1:
				outputChannel <- message
			case message := <-inputChannel2:
				outputChannel <- message
			}
		}
	}()

	return outputChannel
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Got: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
