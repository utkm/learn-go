package main

import (
	"fmt"
	// "strconv"
)

// can declare here but not with := syntax
var z int = 11

var (
	x int  = 1
	y int = 2
)

// upper case => globally visible

func main() {
	// can redeclare vars
	var x int = 10
	y := 10.

	// typecasting:
	var u int
	u = int(y)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", u, u)
	fmt.Println(z)
}

// Summary
// Can't redeclare, but can shadow (declare in package scope, and shadow in function scope)
// All vars must be used
// Visibiltiy: (lower case first letter for package scope; upper case first letter to export; no priv scope)
