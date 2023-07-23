package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for ch := 'a'; ch <= 'e'; ch++ {
		fmt.Printf("%c ", ch)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printNumbers() // Start a new Goroutine for printNumbers()
	printLetters()    // Execute printLetters() concurrently with the Goroutine
}
