package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)

			channel1 <- "channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 3)

			channel2 <- "channel 2"
		}
	}()

	for {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)

		case message2 := <-channel2:
			fmt.Println(message2)
		}
	}
}
