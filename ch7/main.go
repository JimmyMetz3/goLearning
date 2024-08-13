package main

import (
	"fmt"
)

func main() {
	floats := []float64{98, 93, 77, 82, 83}
	ints := []int{1, 2, 7, 4, 5}
	fmt.Println(average(floats))
	fmt.Println(add(ints...))
	fmt.Println(sum(ints))
	fmt.Println(halveNum(1))
	fmt.Println(max(1, 2, 3, 7, 4))
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// PROBLEM
/*sum is a function which takes a slice of numbers
and adds them together. What would its function signature look like in Go?*/

func sum(sliceofNums []int) int {
	var total int = 0
	for _, v := range sliceofNums {
		total += v
	}
	return total
}

// PROBLEM
/*Write a function which takes an integer and halves it and returns true if it was even or false
if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).*/

func halveNum(num int) (int, bool) {
	if num%2 == 0 {
		return (num / 2), true
	} else {
		return (num / 2), false
	}
}

// PROBLEM
/* Write a function with one variadic parameter
that finds the greatest number in a list of numbers */

func max(nums ...int) int {
	max := 0
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	return max
}
