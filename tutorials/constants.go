package main

import (
	"fmt"
	// "math"
)

const (
	_ = iota // write only constant
	a
	b
	c
)

func main() {
	// has to be assignable at compile time (arithmetic, bitwise operations, bitshifting)
	const myConst int = 11
	fmt.Printf("%v, %T\n", myConst, myConst)

	// Enumerated constants
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}