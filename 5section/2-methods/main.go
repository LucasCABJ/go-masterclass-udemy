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

// Value receiver (data is copied)
func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName 
}

func (e *Employee) Deactive() {
	e.IsActive = false
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
	jane := Employee{
		ID: 1,
		FirstName: "Jane",
		LastName: "Doe",
		Position: "Night",
		Salary: 15000,
		IsActive: true,
		JoinedAt: time.Now(),
	}
	fmt.Printf("%+v\n", jane)
	jane.Deactive()
	fmt.Printf("%+v\n", jane)
}

