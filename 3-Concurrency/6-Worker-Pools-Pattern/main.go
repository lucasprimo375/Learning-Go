package main

import "fmt"

func main() {
	tasks := make(chan int, 45) // sequence of numbers to calculate
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results

		fmt.Println(result)
	}
}

// tasks only receives while results only sends
func worker(tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- fibonacci(task)
	}
}

func fibonacci(index int) int {
	if index <= 1 {
		return index
	}

	return fibonacci(index-1) + fibonacci(index-2)
}
