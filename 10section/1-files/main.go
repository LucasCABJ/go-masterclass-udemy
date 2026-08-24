package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// filePath := "10section/1-files/text.txt"
	// data := "welcome to the Go programming language"
	// err := os.WriteFile(filePath, []byte(data), 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// content, err := os.ReadFile(filePath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(content))

	// file2, err := os.Create("10section/1-files/file-via-create.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file2.Close()

	// _, err = file2.WriteString("Welcome all java and python devs!")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	languagesFile, err := os.Open("./10section/1-files/languages.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(languagesFile)
	lineNum := 1
	for scanner.Scan() {
		fmt.Println(lineNum, scanner.Text())
		lineNum++
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		log.Fatal(err)
	}
}
