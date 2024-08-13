package main

import "fmt"

// func that DOES use Pointers
// In Go a pointer is represented using the * (astersk) character followed by the type of the stored value
// * is also used to "dereference" pointer variables. Dereferenceing a pointer gives us access to the value the pointer points to.
func zero(xPtr *int) {
	*xPtr = 0
}

func one(yPtr *int) {
	*yPtr = 1
}

//PROBLEM
/* Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1). */

func swap(x *int, y *int) {
	tempValue := new(int)
	*tempValue = *x
	*x = *y
	*y = *tempValue
}

func main() {
	// POINTERS

	//Pointers reference a location in memory where a value is stored rather than the value itself

	x := 5
	zero(&x) // Use the & operator to find the address of a variable
	// &x returns a *int
	fmt.Println(&x) // returns x location in memory!
	fmt.Println(x)  // x is still 5. Function does not change it's value

	// Another way to get a pointer is to use the built-in new() function

	yPtr := new(int) // new() takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it
	one(yPtr)
	fmt.Println(*yPtr) // y is 1 // If you take the * away, you will just receive the location in memory of yPtr

	// PROBLEM
	// How do you assign a value to a pointer?
	testPointer := new(int)
	*testPointer = 2
	fmt.Println(testPointer)
	fmt.Println(*testPointer)

	testx := 10
	testy := 7
	swap(&testx, &testy)
	fmt.Println(testx, testy)
}
