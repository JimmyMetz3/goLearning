package main

import "fmt"

func main() {
	// Arrays
	/* Arrays are a numbered sequence of elements of
	a single type with a fixed length*/
	var x [5]int
	x[4] = 100

	fmt.Println(x)     // prints the whole array
	fmt.Println(x[4])  // prints a specific value in the array
	fmt.Println(x[1:]) // prints all values starting at index of 1

	// Long way to initialize an Array

	// var floatArray [5]float64
	// floatArray[0] = 98
	// floatArray[1] = 93
	// floatArray[2] = 77
	// floatArray[3] = 82
	// floatArray[4] = 83

	// Short way to initialize an Array in one Line

	// floatArray := [5]float64{98, 93, 77, 82, 83}

	// Another way

	floatArray := [5]float64{
		98,
		93,
		77,
		82,
		83, // This comma is required when using this format
	}

	var floatArrayTotal float64

	// Change a value

	floatArray[2] = 76

	// Unoptimized way of looping through the array with a for loop

	// for i := 0; i < 5; i++ {
	// 	floatArrayTotal += floatArray[i]
	// }

	// Optimized way to looping through an array!!

	// this for loop with the "range" keyword gives index and values
	for _, value := range floatArray { // _ is used because we do not need the index
		floatArrayTotal += value
	}

	fmt.Println(floatArrayTotal / float64(len(floatArray)))

	// Long version
	// var ages [3]int = [3]int{20, 25, 30}

	// Short version
	// var ages = [3]int{20,25,30}

	// Another way
	//names := [4]string{"yoshi", "mario", "peach", "bowser"}
}
