package main

import (
	"fmt"
)

func main() {
	a := 10
	b := a
	fmt.Println(a, b)
	a = 7
	fmt.Println(a, b)

	// Can change this behaviour using pointers
	var c int = 10
	var d *int = &c // addressof operator
	fmt.Println(c, d) // prints the memory location of the data
	fmt.Println(c, *d) // de-referencing operator
	c = 21
	fmt.Println(c, *d)
	*d = 42
	fmt.Println(c, *d)

	// The concept of pointer arithmetic is not present in Go (unlike in C/C++)
	// Can import the "unsafe" package, though
	e := [3]int{1,2,3}
	f := &e[0]
	// g := &e[1] - 2 (CANNOT do this)
	fmt.Printf("%v %p\n", e, f)

	var ms *myStruct
	// fmt.Println(ms) // nil is the zero value for a pointer
	ms = new(myStruct)
	(*ms).foo = 21 
	fmt.Println((*ms).foo)
	// dot operator takes precedence so put brackets to fix order of operation
	// but we don't need this! the compiler understands this: (*ms).foo is equiv to ms.foo
}

type myStruct struct {
	foo int
}

// Summary
// Complex types (ex pointers to structs) are automatically dereferenced for brevity
// Slices and maps contain internal pointers, so copies point to the same underlying data