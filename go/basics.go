package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
func name(varname type) return_type{
	code...
}
*/

/*
Types in GoLang
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

func add(x int, y int) int {
	return x + y
}

func add_floats(x, y float64) float64 {
	return x + y + 0.5
}

func swap_strings(x, y string) (string, string) {
	return y, x
}

// Naming the return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // also called naked return
	// preferably to use in short functions
}

func main() {
	// Declaring variables
	var i int
	var j, q int = 1, 2
	var yes, no bool
	z := 6 // only works inside a function
	// := can't be used outisde funcs
	var somevar, somevar2 = "ok", "no"
	fmt.Println(i, yes, no, j, q, z)
	fmt.Println(somevar, somevar2)

	fmt.Println("Some random number:", rand.Intn(100))
	fmt.Println("Square root of 7 is:", math.Sqrt(7))
	fmt.Println("Pi is:", math.Pi)
	fmt.Println(add(55, 31))
	fmt.Println(add_floats(11.5, 22.0))
	a, b := swap_strings("hello", "bye")
	fmt.Println(a, b)     //prints bye, hello
	fmt.Println(split(5)) // prints 2 and 3
}
