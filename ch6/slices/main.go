package main

import "fmt"

func main() {
	// A Slice is a segment of an array. NO FIXED LENGTH, WOOHOO!!
	// var testSlice []float64

	// x := make([]float64, 5, 10) // 5 is the length, 10 refers to the capacity

	arr := [5]float64{1, 2, 3, 4, 5}
	y := arr[0:4] // A range of an array

	for i, value := range y {
		fmt.Println("Index:", i, "Value:", value)
	}

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	slice3 := make([]int, 2)
	copy(slice3, slice1) // copies slice 1 into slice 3. Only the first two numbers are input since slice3 has a limit of 2
	fmt.Println(slice3)
	fmt.Println(slice1, slice2)

	// FROM A RANDOM YOUTUBE VID
	// Leaving the square brackets empty creates a slice
	// var scores = []int{100, 50, 60}
	// scores[2] = 25 // Change the value of 60 to 25
	// How to append to a slice
	// scores = append(scores, 85)

	// fmt.Println(scores, len(scores))

	// rangeOne := scores[1:3] // The low(1) includes the value. The high(3) does not include the value

	// fmt.Println(rangeOne)
}
