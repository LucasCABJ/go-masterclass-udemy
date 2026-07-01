package main

import "fmt"

type Person interface {
	GetName() string
}

type BusinessPerson struct {
	ID int
	FirstName string
}

type ID int

func (id ID) String() string {
	return fmt.Sprintf("ID: %d", id)
}

func (e BusinessPerson) GetName() string {
	return e.FirstName
}

func (e BusinessPerson) String() string {
	return fmt.Sprintf("BusinessPerson: ID: %d, Name: %s",  e.ID, e.FirstName)
}

func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

func main() {
	jane := BusinessPerson{
		ID: 2,
		FirstName: "Jane",
	}
	fmt.Println(jane)

	id := ID(123)
	fmt.Println(id)
}
