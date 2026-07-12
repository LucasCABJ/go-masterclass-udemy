package main

import (
	"fmt"
)

// channels is a way for corroutines to communicate
func main() {
	messages := make(chan string, 4) // unbuffered

	fmt.Println("Sending messages to buffered channel")
	messages <- "Hello 1"
	messages <- "Hello 2"
	messages <- "Hello 3"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
