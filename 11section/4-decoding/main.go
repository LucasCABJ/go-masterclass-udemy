package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Password string `json:"-"`
	IsActive bool   `json:"is_active"`
}

var payload = `{"name": "Jane","age": 20,"phone": "123-456-789","is_active": true}`

func main() {
	var u user
	dec := json.NewDecoder(strings.NewReader(payload))
	if err := dec.Decode(&u); err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}