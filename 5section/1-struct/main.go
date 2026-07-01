package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	FirstName string
	LastName string
	Position string
	Salary int
	IsActive bool
	JoinedAt time.Time
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName 
}

func NewEmployee(id int, firstName, lastName, position string, salary int) Employee {
	return Employee{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		Position: position,
		Salary: salary,
		IsActive: true,
		JoinedAt: time.Now(),
	}
}

func main() {

	jane := Employee {
		ID: 1,
		FirstName: "Jane",
		LastName: "Doe",
		Position: "Software Engineer",
		Salary: 80000,
		IsActive: true,
		JoinedAt: time.Now(),
	}
	fmt.Printf("%+v\n", jane)
	fmt.Println(jane.ID)
	fmt.Println(jane.FirstName)
	fmt.Println(jane.LastName)
	fmt.Println(jane.Position)
	fmt.Println(jane.Salary)
	fmt.Println(jane.IsActive)
	fmt.Println(jane.JoinedAt)

	joe := NewEmployee(2, "Joe", "Smith", "Product Manager", 90000)
	fmt.Printf("%+v\n", joe)
	fmt.Println(joe.ID)
	fmt.Println(joe.FirstName)
	fmt.Println(joe.LastName)
	fmt.Println(joe.Position)
	fmt.Println(joe.Salary)
	fmt.Println(joe.IsActive)
	fmt.Println(joe.JoinedAt)

	joePtr := &joe
	joePtr.Salary = 95000
	fmt.Println(joePtr)
	fmt.Println(joePtr.Salary)

	fmt.Println(joePtr.FullName())

}
