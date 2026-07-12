package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("sayHello", message)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(5)


	fmt.Println("Hello from Main() Goroutine")

	go sayHello("Hello World!", time.Second, &wg)
	go sayHello("Hello World 2!", time.Second, &wg)
	go sayHello("Hello World! from 2 secs", 2*time.Second, &wg)
	go sayHello("Hello World! from 5 secs", 5*time.Second, &wg)
	go sayHello("Hello World! from 1 secs", 1*time.Second, &wg)
	
	fmt.Println("Last message from Main() Goroutine")

	wg.Wait()
}
