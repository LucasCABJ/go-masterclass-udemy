package main

import (
	"errors"
	"fmt"
	"strings"
)

func divide(a, b int) (int, error) { // by convetion the error return value is the last one in the list of multiple return values
	if b == 0 {
		return 0, errors.New("Can't divide by 0")
	}
	return a / b, nil
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]
	// return firstName, lastName (DONT NEED)
	return // dont need to declare firstName and lastName here because they are already declared as specified as return values
}

func main() {

	result, err := divide(4, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	firstName, lastName := splitName("Lucas Caraballo")
	fmt.Println(firstName)
	fmt.Println(lastName)
}
