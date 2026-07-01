package main

import "fmt"

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modifyValue: %+v\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		fmt.Println("val is nil")
		return
	}
	*val = *val * 10
	fmt.Printf("modifyPointer: %+v\n", *val)
}

func main() {
	// pointers give us direct access to memory. That's nice.
	// pointers in go doesn't support arithmetic.
	// they save memory address of variables
	// age := 10
	// agePtr := &age
	// fmt.Printf("age: %d\n", agePtr)
	// fmt.Printf("age: %d\n", &age)
	num := 10
	modifyValue(num)
	fmt.Printf("num: %+v\n", num)
	modifyPointer(&num)
	fmt.Printf("num: %+v\n", num)

	grade := 50
	gradePtr := &grade
	fmt.Printf("Gradeptr grade: %+v\n", gradePtr)
	fmt.Printf("GradePtr: %+v\n", gradePtr)
}