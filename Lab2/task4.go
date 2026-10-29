// Lab 4

//go:build ignore

package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func mod(a int, b int) int {
	return a % b
}

func main() {
	var choice int
	var x int
	var y int

	for {
		fmt.Println("******* Mini Calculator *******")
		fmt.Println("1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit")
		fmt.Print("Choose: ")
		fmt.Scanln(&choice)

		// exit
		if choice == 6 {
			fmt.Println("Exit")
			break
		}

		// invalid choice
		if choice < 1 || choice > 6 {
			fmt.Println("Error: invalid operator")
			continue
		}

		fmt.Print("Enter a: ")
		fmt.Scanln(&x)
		fmt.Print("Enter b: ")
		fmt.Scanln(&y)

		switch choice {
		case 1:
			fmt.Println("Result:", add(x, y))
		case 2:
			fmt.Println("Result:", sub(x, y))
		case 3:
			fmt.Println("Result:", mul(x, y))
		case 4:
			if y == 0 {
				fmt.Println("Error: division by zero")
			} else {
				fmt.Println("Result:", div(x, y))
			}
		case 5:
			if y == 0 {
				fmt.Println("Error: modulo by zero")
			} else {
				fmt.Println("Result:", mod(x, y))
			}
		}

		fmt.Println()
	}
}