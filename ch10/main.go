package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	// CONCURRENCY
	/* A goroutine is a function that is capable of running
	concurrently with other functions*/
	// A goroutine will immediately return as soon as it is finished
	// go f(0) // This finishes almost instantly
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input) // Input this Scanln so the application doesn't stop instantly
}
