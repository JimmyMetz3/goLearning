package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for ind, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println("Index of: " , os.Args[ind] , " is " , ind)
	}
	fmt.Println(s)
}
