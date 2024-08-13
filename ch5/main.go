package main

import "fmt"

func main() {
	// i := 1 declared i in the first portion of the for loop
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i, "Even")
	// 	} else {
	// 		fmt.Println(i, "Odd")
	// 	}
	// switch i {
	// case 0:
	// 	fmt.Print("Zero")
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 4:
	// 	fmt.Println("Four")
	// case 5:
	// 	fmt.Println("Five")
	// case 6:
	// 	fmt.Println("Six")
	// case 7:
	// 	fmt.Println("Seven")
	// case 8:
	// 	fmt.Println("Eight")
	// case 9:
	// 	fmt.Println("Nine")
	// case 10:
	// 	fmt.Println("Ten")
	// default:
	// 	fmt.Println("Unknown Number")
	// }
	// }

	// PROBLEM
	/* Write a program that prints out all the numbers
	evenly divisible by 3 between 1 and 100. (3, 6, 9,
	etc.) */

	// // Loop through nums 1 - 100
	// for i := 1; i <= 100; i++ {
	// 	//Check to see if the number is evenly divisible by 3
	// 	if i%3 == 0 {
	// 		fmt.Println(i, "is evenly divisible by 3")
	// 	} else {
	// 		fmt.Println(i, "is not evenly divisible by 3")
	// 	}
	// }

	//PROBLEM
	/* Write a program that prints the numbers from 1
	to 100.
	But for multiples of three print "Fizz" instead of the number
	and for the multiples of fiveprint "Buzz".
	For numbers which are multiples
	of both three and five print "FizzBuzz".*/

	// Print numbers 1 - 100
	for i := 1; i <= 100; i++ {
		//If the number is divisible by 3 and 5 print fizzbuzz
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i, "Whomp Whomp")
		}

	}
}
