package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond) // Simulate some work by the producer
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Consuming %d\n", num)
		time.Sleep(1 * time.Second) // Simulate some work by the consumer
	}
}

func main() {
	normalCh := make(chan int)      // Normal channel
	bufferedCh := make(chan int, 3) // Buffered channel with capacity 3

	fmt.Println("Using Normal Channel:")
	go producer(normalCh)
	consumer(normalCh)

	time.Sleep(2 * time.Second) // Wait for the producer to finish

	fmt.Println("\nUsing Buffered Channel:")
	go producer(bufferedCh)
	consumer(bufferedCh)
}
