package main

import "fmt"

func main() {
	username := "山月日"
	fmt.Println(len(username)) // out: 9, returns len of bytes
	for _, v := range username {
		fmt.Printf("%c\n", v) // v is rune type, rune is go unicode char type
	}
}
