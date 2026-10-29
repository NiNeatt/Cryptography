// Lab 5

//go:build ignore

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var text string

	fmt.Print("Enter a text: ")
	fmt.Scanln(&text)

	// In binary 
	fmt.Print("Binary: ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%08b ", text[i])
	}
	fmt.Println()

	// hexadecimal

	fmt.Printf("Hex: %x\n", text)

	// base64
	fmt.Println("Base64:", base64.StdEncoding.EncodeToString([]byte(text)))
}