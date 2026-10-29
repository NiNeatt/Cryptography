// Lab 3


//go:build ignore

package main

import "fmt"

func myAND(a int, b int) int {
	return a & b
}

func myOR(a int, b int) int {
	return a | b
}

func myXor(a int, b int) int {
	return a ^ b
}

func myNOT(a int) int {
	return ^a
}

func myShiftLeft(a int, c int) int {
	return a << c
}

func myShiftRight(a int, c int) int {
	return a >> c
}

func main () {
	var x int
	var y int
	var c int
	var answer int
	fmt.Print("Please Enter the Value of x: ")
	fmt.Scanln(&x)
	fmt.Print("Please Enter the Value of y: ")
	fmt.Scanln(&y)

	fmt.Printf("x = %d (binary %b)\n", x, x)

	fmt.Printf("y = %d (binary %b)\n", y, y)

	fmt.Print("How many places to shift? ")
	fmt.Scanln(&c)


	answer = myAND(x, y)
	fmt.Printf("AND Answer = %d (binary %b)\n", answer, answer)

	answer = myOR(x, y)
	fmt.Printf("OR  Answer = %d (binary %b)\n", answer, answer)

	answer = myXor(x, y)
	fmt.Printf("XOR Answer = %d (binary %b)\n", answer, answer)

	answer = myNOT(x)
	fmt.Printf("NOT x   = %d (binary %b)\n", answer, answer)

	answer = myNOT(y)
	fmt.Printf("NOT y   = %d (binary %b)\n", answer, answer)


	answer = myShiftLeft(x, c)
	fmt.Printf("Shift left  = %d (binary %b)\n", answer, answer)

	answer = myShiftRight(x, c)
	fmt.Printf("Shift right = %d (binary %b)\n", answer, answer)
	
}

