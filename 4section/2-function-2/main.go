package main

import "fmt"

func factorial(n int) int { // RECURSIVE FUNC
	if n <= 1 { // base case
		return 1
	}
	return n * factorial(n-1) // recursive case
}

func intSeq() func() int { // CLOSURE FUNC
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(factorial(5))

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	logger := func(msg string) {
		fmt.Println(msg)
	}

	logger("test")
	logger("test2")
}
