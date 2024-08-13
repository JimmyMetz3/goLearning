package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

func circleArea(cir *Circle) float64 {
	return math.Pi * cir.r * cir.r
}

// Method

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// A struct is a type which contains named fields

// Example struct

type Circle struct { // type keyword introduces a new type
	//x float64 // each field has a name and type
	//y float64
	//r float64
	// We can collapse fields that have the same type
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

// func (r *Rectangle) area() float64 {
// 	l := distance(r.x1, r.y1, r.x1, r.y2)
// 	w := distance(r.x1, r.y1, r.x2, r.y1)
// 	return l * w
// }

// Struct's fields usually represent a "has-a" relationship

type Person struct {
	Name string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person // This represents a "Is-a" relationship. An Android Is-a person... right?
	Model  string
}

func main() {
	// STRUCTS
	// We can create instances of the Circle type a few diff ways

	//var c Circle
	// by default, all values for our circle will be 0s

	// We can use the new function as well
	/* Using the new() func allocates memory for all fields,
	sets each of them to their zero value and returns a pointer*/
	// c := new(Circle)
	// Create a new struct and input values
	// c := Circle{x: 0,y: 0,r: 5}
	// OR
	c := Circle{0, 0, 5}

	// c.x = 0
	// c.y = 0
	// c.r = 5
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(circleArea(&c))
	fmt.Println(c.area())
	pers := Person{Name: "Jimmy"}
	pers.talk()

	andr := new(Android)
	// Real Path
	//andr.Person.Name = "Android17"
	// Shorthand Path
	andr.Name = "Android18"

	// Long path
	//andr.Person.talk()
	// Shorthand
	andr.talk()

	android20 := Android{}
	android20.Name = "Android20"
	android20.talk()
}
