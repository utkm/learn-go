package main

import (
	"fmt"
)

func main() {
	grades := [...]int{100, 98, 92}
	var students [3]string
	fmt.Printf("Grades: %v\n", grades)
	students[0] = "John"
	students[1] = "Jeff"
	students[2] = "Bob"
	fmt.Printf("Students: %v\n", students)
	fmt.Println(len(students))

	// Arrays are considered values
	// Slices are reference types
	a := []int{1, 2, 3} // slice
	// b := &a // address of operations vs copy of a
	// b[1] = 5
	fmt.Println(a)
	fmt.Printf("Capacity: %v\n", cap(a))
	// fmt.Println(b)

	c := make([]int, 3) // type, length, capacity
	fmt.Println(c)
	fmt.Printf("Length: %v\n", len(c))
	fmt.Printf("Capacity: %v\n", cap(c))
	c = append(c, 1)
	fmt.Println(c)
	fmt.Printf("Length: %v\n", len(c))
	fmt.Printf("Capacity: %v\n", cap(c)) // copies existing elements to a new array

	// Removing element from middle
	x := []int{1, 2, 3, 4, 5}
	y := append(x[:2], x[3:]...)
	fmt.Println(y)
	fmt.Println(x)
}

// Summary
// Arrays are collections of items with the same type, and have fixed size
// Functions: len, cap (length of underlying array), append, 
// Copies refer to underlying data