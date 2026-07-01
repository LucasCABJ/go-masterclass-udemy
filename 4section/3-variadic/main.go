package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func config(numbers ...int) { 
	if (len(numbers) > 0) {
		first := numbers[0]
		fmt.Println("First number: ", first)
		return
	}
	fmt.Println("Default number")
}

func main() {
	fmt.Println(1, 2, 3, 4, 5)
	config()
	config(23)
}
