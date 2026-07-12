package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Welcome to the wonderful world of Go!"

	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println("Encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString("V2VsY29tZSB0byB0aGUgd29uZGVyZnVsIHdvcmxkIG9mIEdvIQ==")
	if err != nil {
		fmt.Println("Error decoding string")
		return
	}
	fmt.Println("Decoded: ", string(decoded))
}
