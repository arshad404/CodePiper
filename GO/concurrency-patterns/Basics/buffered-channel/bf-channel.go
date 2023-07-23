package main

import "fmt"

func main() {
	ch := make(chan int, 3) // Create a buffered channel with a capacity of 3

	ch <- 1 // Send data to the channel
	ch <- 2 // Send more data
	ch <- 3 // Send even more data

	// ch <- 4 // Sending a 4th value would cause a deadlock because the channel is full

	fmt.Println(<-ch) // Receive data from the channel
	fmt.Println(<-ch) // Receive more data
	fmt.Println(<-ch) // Receive even more data

	// fmt.Println(<-ch) // Receiving a 4th value would cause a deadlock because the channel is empty
}
