package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup that this Goroutine is done when the function returns
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup that this Goroutine is done when the function returns
	for ch := 'a'; ch <= 'e'; ch++ {
		fmt.Printf("%c ", ch)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Add the number of Goroutines to wait for

	go printNumbers(&wg) // Start a new Goroutine for printNumbers()
	go printLetters(&wg) // Start a new Goroutine for printLetters()

	wg.Wait() // Wait until all Goroutines in the WaitGroup finish
	fmt.Print("Okay you can go")
}
