package main

import "fmt"

type Person interface {
	GetName() string
}

type Employee struct {
	ID int
	FirstName string
}

type BusinessPerson struct {
	ID int
	FirstName string
}

func (e Employee) GetName() string {
	return e.FirstName
}

func (e BusinessPerson) GetName() string {
	return e.FirstName
}

func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

func main() {
	// joe := Employee{
	// 	ID: 1,
	// 	FirstName: "Joe",
	// }
	jane := BusinessPerson{
		ID: 2,
		FirstName: "Jane",
	}

	fmt.Println(jane)
	// displayPerson(joe)
	// displayPerson(jane)
}
