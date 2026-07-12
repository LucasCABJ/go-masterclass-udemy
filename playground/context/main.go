package main

import (
	"context"
	"fmt"
	"time"
)

// context is a tool that carries values, deadlines (timeouts) and signals across out go program
// typical uses: controlling timeouts, cancelling goroutines, passing metadata across application
func main() {
	ctx := context.Background()
	exampleTimeout(ctx)

	exampleWithValues()
}

func exampleTimeout(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	done := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Called the API")
	case <-ctxWithTimeout.Done():
		fmt.Println("Oh no my timeout expired")
		// do some logic
	}
}

func exampleWithValues() {
	type key int
	const Userkey key = 0

	ctx := context.Background()
	ctxWithValue := context.WithValue(ctx, Userkey, "123")

	if userId, ok := ctxWithValue.Value(Userkey).(string); ok {
		fmt.Println("this is the userID:", userId)
	} else {
		fmt.Println("No user found")
	}
}
