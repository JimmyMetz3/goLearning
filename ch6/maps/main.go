package main

import "fmt"

func main() {
	// A map is an unordered collection of key-value pairs. Also known as an associative array, a hash table, or a dictionary

	// This will cause a runtime error

	// var x map[string]int

	// Long way of making maps
	// names := make(map[string]int)
	// names["Jimmy"] = 30
	// names["Cayla"] = 24
	// names["Levi"] = 30
	// names["Cullen"] = 33

	// Short way of making maps

	names := map[string]int{
		"Jimmy":  30,
		"Cayla":  24,
		"Levi":   30,
		"Cullen": 33,
	}

	fmt.Println(names["Jimmy"]) // Keys are case sensitive!!!
	fmt.Println(len(names))

	delete(names, "Jimmy")
	fmt.Println(len(names))

	// fmt.Println(names["Rick"]) // This returns 0 aka the empty value of an int

	// This returns 1: the outcome of the lookup = 0 AND a boolean to show if the lookup was successful
	// name, ok := names["Rick"]
	// fmt.Println(name, ok)

	if name, ok := names["Rick"]; ok {
		fmt.Println(name, ok)
	}

	// You can nest Maps inside of each other

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}

	fmt.Println(elements["H"])
	// for i, v := range elements {
	// 	fmt.Println(i, v)
	// }
	y := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(y[2:5])

	// PROBLEM - WRITE A PROGRAM THAT FINDS THE SMALLEST NUMBER IN THIS LIST

	numbers := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var lowestNum int = 100
	for _, value := range numbers {
		if value < lowestNum {
			lowestNum = value
		}
	}
	fmt.Println(lowestNum)
}
