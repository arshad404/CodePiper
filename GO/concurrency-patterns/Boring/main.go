package main

import (
	"fmt"
	"sync"
)

// Boring function that returns a channel to send boring messages
func boring(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch) // Close the channel when the goroutine exits
		for i := 1; i <= 5; i++ {
			ch <- fmt.Sprintf("%s: %d", msg, i)
		}
	}()
	return ch
}

func main() {
	aliceCh := boring("Alice")
	bobCh := boring("Bob")

	// Receive messages concurrently
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for msg := range aliceCh {
			fmt.Println(msg)
		}
		wg.Done()
	}()

	go func() {
		for msg := range bobCh {
			fmt.Println(msg)
		}
		wg.Done()
	}()

	wg.Wait()
}
