// LAb 6

//go:build ignore

package main

import "fmt"

func xorEncrypt(text string, key byte) string {

	data := []byte(text)
	for i := 0; i < len(data); i++ {
		data[i] = data[i] ^ key
	}
	return string(data)
}

func xorRepeat(text string, key string) string {
	data := []byte(text)

	for i := 0; i < len(data); i++ {
		data[i] = data[i] ^ key[i%len(key)]
	}

	return string(data)
}

func main() {
	var text string
	var choice int
	var cipher string
	var plain string

	fmt.Print("Enter a message: ")
	fmt.Scanln(&text)

	fmt.Println("1) Single byte key (a number)")
	fmt.Println("2) Repeating key (a word)")
	fmt.Print("Please Choose One of Your Choices: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var keyNumber int

		fmt.Print("Enter a key number from 0 to 255: ")
		fmt.Scanln(&keyNumber)

		if keyNumber < 0 || keyNumber > 255 {
			fmt.Println("Error: key must be from 0 to 255")
			return
		}

		key := byte(keyNumber)

		// encrypt
		cipher = xorEncrypt(text, key)

		// decrypt
		plain = xorEncrypt(cipher, key)

	case 2:
		var key string

		fmt.Print("Enter a key word: ")
		fmt.Scanln(&key)

		if len(key) == 0 {
			fmt.Println("Error: key cannot be empty")
			return
		}

		// encrypt
		cipher = xorRepeat(text, key)

		// decrypt
		plain = xorRepeat(cipher, key)

	default:
		fmt.Println("Error: invalid choice")
		return
	}

	fmt.Println("Original message :", text)
	fmt.Printf("Ciphertext        :%x\n",cipher)
	fmt.Println("Decrypted message:", plain)
}