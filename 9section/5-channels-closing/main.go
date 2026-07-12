package main

import (
	"fmt"
	"time"
)

// channels is a way for corroutines to communicate
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			r, ok := <- jobs
			if !ok {
				fmt.Println("Channel closed")
				done <- true
				return
			}
			fmt.Println("Message got: ", r)
		}
	}()

	for i := 1; i <= 6; i++ {
		jobs <- i
		fmt.Println("Sent: ", i)
		time.Sleep(time.Second)
	}

	close(jobs)
	<-done 
}
