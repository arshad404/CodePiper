package main

import (
	"fmt"
	"time"
)

func taskOne(resultChan chan<- string) {
	for {
		// Simulating some task
		time.Sleep(2 * time.Second)
		resultChan <- "Task One Completed"
	}
}

func taskTwo(resultChan chan<- string) {
	for {
		// Simulating another task
		time.Sleep(3 * time.Second)
		resultChan <- "Task Two Completed"
	}
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	resultChan := make(chan string)

	// Launching the goroutines for taskOne and taskTwo
	go taskOne(resultChan)
	go taskTwo(resultChan)

	// Main goroutine with select statement
	for {
		select {
		case <-ticker.C:
			// Perform the periodic task here
			fmt.Println("Tick")

		case result := <-resultChan:
			// Handling the results from the goroutines
			fmt.Println(result)
		}
	}
}
