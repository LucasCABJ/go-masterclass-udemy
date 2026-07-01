package main

import "fmt"

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: Deferred")
	fmt.Println("Function simpleDefer: Middle")
}

func lifoSimpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: 1 Deferred")
	defer fmt.Println("Function simpleDefer: 2 Deferred")
	defer fmt.Println("Function simpleDefer: 3 Deferred")
	fmt.Println("Function simpleDefer: Middle")
}

func main() {
	defer func() {
		fmt.Println("Before return in main!")
	}()
	// simpleDefer()
	lifoSimpleDefer()
	fmt.Println("Last in main")
}
