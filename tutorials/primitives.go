// Bool, Int, Float, Complex, 

package main

import (
	"fmt"
)

func main() {
	// variable initializations have a 0 value
	var n bool
	fmt.Printf("%v, %T\n", n, n)

	// Numericals
	// unsigned is non-negative
	a := 10 // 1010
	b := 3  // 0011
	fmt.Println(a / b)
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100
	c := 8 // 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0 = 1

	var w complex64 = 1 + 2i
	var z complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", w, w)
	fmt.Printf("%v, %T\n", z, z)


	// Text Type
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
}