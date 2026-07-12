package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}

func main() {

	jane := user{
		Name:     "Jane",
		Age:      20,
		IsActive: true,
		Phone:    "123-456-789",
	}

	byteSlice, err := json.Marshal(jane)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteSlice))
}
