package main

import "fmt"

const greeting string = "Hello!" //Allows all funcs to access this var

func main() {

	var x string
	x = "first "
	var y = "third"
	fmt.Println(x)
	// x = x + "second"
	x += "second" // append string to x
	fmt.Println(x)
	fmt.Println(x == y) //false

	// WAYS TO ASSIGN VARIABLES

	// var x = "Hi"
	// x := "Hi"
	// var x string = "Hi"

	fmt.Println(greeting)
	greet()

	// ASSIGN MULTIPLE VARIABLES
	var ( //could use const too!
		a int    = 5
		b string = "Hi"
		c        = 9
	)
	fmt.Println(a, b, c)
}

func greet() {
	fmt.Println(greeting)
}
