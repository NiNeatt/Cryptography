// Lab 2


//go:build ignore
package main 

import "fmt"

func main () {
	var x int
	var y int

	fmt.Print("Please Enter the vaalue of x: ")
	// get the  value of integer1 from user input
	fmt.Scanln(&x)
	// get the value of inter2 from user 2
	fmt.Print("Please Enter the value of y: ")
	fmt.Scanln(&y)

	var result bool

	if (x > 0 && y>0) {
		result = true
		fmt.Println("Both positive (&&):",result)
	}else {
		result =false
		fmt.Println("Both positive (&&)",result)
	}

	if (x>y || y> x) {
		result = true
		fmt.Println("One greater than the other (||):",result)
	}else {
		result =false
		fmt.Println("One greater than the other (||):",result)
	}

	if (x != y) {
		result = true
		fmt.Println("Not equal (!):",result)
	}else {
		result =false
		fmt.Println("Not equal (!):",result)
	}


}