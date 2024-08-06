package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println(strings.Join(os.Args[1:], " "))
	// Shortest line without caring about formatting
	fmt.Println(os.Args[1:])
}
