package main

import "fmt"

func main() {
	// array allocates contiguous memory
	var numbers [2]int
	fmt.Printf("%+v\n", numbers)
	numbers[0] = 1
	numbers[1] = 2
	fmt.Printf("%+v\n", numbers)

	// short hand initialization
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Printf("%+v\n", primes)

	// len() to iterate with c style loop
	for i := 0; i < len(primes); i++ {
		fmt.Printf("%d\n", primes[i])
	}

	// matrix
	var matrix [2][3]int
	matrix[0][1] = 100
	fmt.Printf("%+v\n", matrix)
}
