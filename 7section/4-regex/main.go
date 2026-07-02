package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	text1 := "Hello World! Welcome to GO"
	regGo, err := regexp.Compile(`GO`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("text: %s, matches '%s': %t\n", text1, regGo.String(), regGo.MatchString(text1))

	regWords, err := regexp.Compile("[a-zA-Z]+")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	words := regWords.FindAllString(text1, -1)
	fmt.Printf("%T, %+v\n", words, words)
}
