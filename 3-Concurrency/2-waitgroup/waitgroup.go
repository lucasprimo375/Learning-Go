package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(4) // starts counter with 2, which means wait for 2 goroutines

	go func() {
		write("Hello World!")
		waitgroup.Done() // removes one from the counter; -1
	}()

	go func() {
		write("Programming with Go")
		waitgroup.Done()
	}()

	go func() {
		write("Good morning!")
		waitgroup.Done()
	}()

	go func() {
		write("Hello there!")
		waitgroup.Done()
	}()

	waitgroup.Wait() // wait for counter to reach zero
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
