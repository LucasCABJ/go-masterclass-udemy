package main

import (
	"fmt"
)

func main() {

	// C - style loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// While stile
	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}
	fmt.Println("---- INFINITE LOOP ----")

	// Infinite loop
	counter := 0
	for {
		counter++
		if counter >= 5 {
			break
		}
		fmt.Println(counter)
	}

	fmt.Println("---- SKIPPING ----")
	// Skipping
	for i := 0; i < 10; i++ {
		if (i % 2 == 0) {
			continue
		}
		fmt.Println(i)
	}

	// array foreach
	fmt.Println("---- ARRAY (Foreach) ----")
	items := [3]string{"Python", "C", "Java"}
	for i, v := range items {
		fmt.Printf("%d %s\n", i, v)
	}

	
	// skipping index in foreach
	fmt.Println("---- ARRAY (Foreach) skip index ----")
	for _, v := range items {
		fmt.Printf("%s\n",  v)
	}

	// skipping index in foreach
	fmt.Println("---- ARRAY (Foreach) skip value ----")
	for i := range items {
		fmt.Printf("%d\n",  i)
	}
}
