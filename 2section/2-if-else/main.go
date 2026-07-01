package main

import "fmt"

func main() {
	tmp := 25
	if tmp > 30 {
		fmt.Println("Greater than 30")
	} else {
		fmt.Println("temperature is less than 30")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Failed")
	}

	userAccess := map[string]bool {
		"jane": true,
		"john": false,
	}
	hasAccessInput := "john"

	if hasAccess, ok := userAccess[hasAccessInput]; ok && hasAccess {
		fmt.Printf("%s has access!\n", hasAccessInput)
	} else {
		fmt.Printf("%s doesn't have access!\n", hasAccessInput)
	}
}
