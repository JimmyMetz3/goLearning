package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	//*************STRINGS**************
	// The built-in strings package provides string manipulation!

	// true
	fmt.Println(strings.Contains("Hello", "ell"))

	// 2 "t"s are located in the word test
	fmt.Println(strings.Count("test", "t"))

	// true
	fmt.Println(strings.HasPrefix("test", "te"))

	// true
	fmt.Println(strings.HasSuffix("test", "st"))

	// 1
	fmt.Println(strings.Index("test", "e"))

	// "a-b"
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))

	//"aaaaa"
	fmt.Println(strings.Repeat("a", 5))

	// "bbaa"
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))

	// []string{"a","b","c","d","e"}
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	// test
	fmt.Println(strings.ToLower("test"))

	// "TEST"
	fmt.Println(strings.ToUpper("test"))

	// Convert slice of string into bytes
	arr := []byte("test")
	fmt.Println(arr)

	// Convert bytes into a string
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)

	//*************FILES**************

	// CREATE - to create a file, use os.Create()
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	return
	// }
	// defer file.Close()

	// file.WriteString("test")

	// bs, err := os.ReadFile("test.text")
	// if err != nil {
	// 	return
	// }
	// stri := string(bs)
	// fmt.Println(stri)

	//*************ERRORS**************
	erro := errors.New("Error message")
	fmt.Println(erro)
}
