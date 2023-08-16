package main

import (
	"fmt"
	"sync"
)

func producer(id int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- id*10 + i
		}
	}()
	return ch
}

func fanIn(inputs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	output := make(chan int)

	// Function to copy values from a channel to the output channel
	copy := func(c <-chan int) {
		defer wg.Done()
		for val := range c {
			output <- val
		}
	}

	wg.Add(len(inputs))
	for _, ch := range inputs {
		go copy(ch)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func main() {
	// Step-1 Producer
	ch1 := producer(1)
	ch2 := producer(20)
	ch3 := producer(300)

	// Step-2 FanIn
	mergedOutputChannel := fanIn(ch1, ch2, ch3)

	// Step-3 Consumer
	for val := range mergedOutputChannel {
		fmt.Println(val)
	}
}
