// Lab 1

//go:build ignore

package main

import "fmt"

func main() {
	var integer1 int
	var integer2 int

	fmt.Print("Please Enter Integer1: ")
	// get the  value of integer1 from user input
	fmt.Scanln(&integer1)
	// get the value of inter2 from user 2
	fmt.Print("Please Enter Integer2: ")
	fmt.Scanln(&integer2)

	original := integer1
	var result int
	// = operator

	integer1 = integer2
	result = integer1
	fmt.Println("integer1 = integer2  => result:", result)

	// += operator
	integer1 = original
	integer1 += integer2
	result = integer1
	fmt.Println("integer1 += integer2 => result:", result)

	// -= operator
	integer1 = original
	integer1 -= integer2
	result = integer1
	fmt.Println("integer1 -= integer2 => result:", result)

	// *= operator
	integer1 = original
	integer1 *= integer2
	result = integer1
	fmt.Println("integer1 *= integer2 => result:", result)

	// /= and %= operator

	if integer2 == 0 {
		fmt.Println("Cannot divide by zero, so /= and %= are skipped")
	} else {
		integer1 = original
		integer1 /= integer2
		result = integer1
		fmt.Println("integer1 /= integer2 => result:", result)

		integer1 = original
		integer1 %= integer2
		result = integer1
		fmt.Println("integer1 %= integer2 => result:", result)
	}

}
