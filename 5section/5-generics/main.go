package main

import "fmt"

type Number interface {
	int | float64 | float32 | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8
}

func Sum[T Number](numbers ...T) T {
	var total T
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(Sum(1.5, 2.5, 3.5))
}
