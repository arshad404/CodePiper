package main

import (
	"fmt"
)

func sender(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i // Send data to the channel
	}
	close(ch) // Close the channel after sending all data
}

func receiver(ch <-chan int) {
	for num := range ch {
		fmt.Printf("%d ", num) // Receive data from the channel
	}
}

func main() {
	ch := make(chan int)

	go sender(ch) // Start a new Goroutine for the sender
	receiver(ch)  // Execute the receiver concurrently with the Goroutine
}
