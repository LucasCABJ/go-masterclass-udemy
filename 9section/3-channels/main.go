package main

import (
	"fmt"
	"time"
)

// channels is a way for corroutines to communicate

func main() {
	messages := make(chan string) // unbuffered
	go func() {
		fmt.Println("Sending message to messages channel")
		messages <- "Hello to messages channel!"
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("About to get a message from channel")
	msg := <-messages
	fmt.Println(msg)

	msg = <-messages
	fmt.Println(msg)
}
