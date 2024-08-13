package main

import "fmt"

func main() {
	// Prompt for an input
	// fmt.Print("Enter a number: ")
	// var input float64
	// fmt.Scanf("%f", &input)

	// Multiply the input by 2
	// output := input * 2

	// Print the result
	// fmt.Println(output)

	// PROBLEMS

	//1. What are two ways to create a new variable?

	//ANSWER
	// var num int = 5
	// num := 5
	// var num = 5

	/*2. What is the value of x after running:
	x := 5; x += 1?*/

	//ANSWER
	// 6
	// testnum := 5
	// testnum += 1
	// fmt.Println(testnum)

	/*3. What is scope and how do you determine the
	scope of a variable in Go?*/

	//vars can only be accessed within each block unless it is assigned outside of all variables

	//4. What is the difference between var and const?

	// var can be reassigned, const cannot

	/*5. Using the example program as a starting point,
	write a program that converts from Fahrenheit
	into Celsius. (C = (F - 32) * 5/9)*/

	// take input
	// fmt.Print("Enter the temperature in Fahrenheit: ")
	// var inputTemp float64
	// fmt.Scanf("%f", &inputTemp)

	// calculate the temp to celsius

	// celsiusResult := ((inputTemp - 32) * 5 / 9)

	// fmt.Println("That is", celsiusResult, "in celsius")

	/*6. Write another program that converts from feet
	into meters. (1 ft = 0.3048 m)*/

	// Prompt for input
	fmt.Print("Enter num of feet to be converted to meters: ")
	var inputFeet float64
	fmt.Scanf("%f", &inputFeet)

	//Convert the measurement to meters

	feetToMeters := (inputFeet * .3048)

	// Output

	fmt.Println("That is", feetToMeters, "in meters.")
}
