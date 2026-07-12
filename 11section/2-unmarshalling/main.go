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

var payload = `
{
 "name": "Jane",
 "age": 20,
 "phone": "123-456-789",
 "is_active": true
}
`

func main() {
	var user user
	err := json.Unmarshal([]byte(payload), &user)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%+v\n", user)
}
