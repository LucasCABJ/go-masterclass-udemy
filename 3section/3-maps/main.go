package main

import "fmt"

func main() {
	studentGrades := map[string]int {
		"Alice": 90,
		"James": 85,
		"Dan": 60,
	}
	fmt.Printf("%+v\n", studentGrades)
	studentGrades["Alice"] = 95
	fmt.Printf("%+v\n", studentGrades)

	if alice, ok := studentGrades["Alice"]; ok {
		fmt.Printf("Alice found! note: %+v\n", alice)
	}
	
	key := "Dan"
	if value, ok := studentGrades[key]; ok {
		fmt.Printf("%s found! note: %+v\n", key, value)
	}

	fmt.Printf("%+v\n", studentGrades)
	delete(studentGrades, "Alice")
	fmt.Printf("%+v\n", studentGrades)

	configs := map[string]int{}
	fmt.Printf("%+v, %T\n", configs, configs)
}