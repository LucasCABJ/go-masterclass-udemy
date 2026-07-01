package main

import (
	"fmt"
	"time"
)

func main() {
	day := "Sunday"
	fmt.Println("Today is ", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend! No work!")
	case "Monday", "Tuesday":
		fmt.Println("Work days. Lots of meetings!")
	default:
		fmt.Println("Mid-week")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good Night!")

	}

}
