package main

import (
	"encoding/json"
	"log"
	"os"
)

type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Password string `json:"-"`
	IsActive bool   `json:"is_active"`
}

var payload = `{
	name: "Jane",
	age: 20,
	phone: "123-456-789",
	is_active: true	
}`

func main() {
	u := user{
		Name: "John Smith",
		Age: 30,
		Phone: "123-456-789",
		IsActive: true,
	}

	dec := json.NewEncoder(os.Stdout)
	if err := dec.Encode(u); err != nil {
		log.Fatal(err)
	}

}