package main

import (
	"fmt"
)

func main() {
	provincePopulations := map[string]int{
		"Ontario":          14570000,
		"Quebec":           8485000,
		"British Columbia": 5071000,
	}

	// interrogate the map and then generate a boolean result
	if pop, ok := provincePopulations["Ontario"]; ok {
		// variables are only defined in the scope of the if statement
		fmt.Println(pop)
	} else if true {
		fmt.Println("For sure")
	} else {
		fmt.Println("Never runs")
	}

	// simple way to compare one variable to multiple possible cases
	switch 3 {
	case 1, 3, 5:
		fmt.Println("one, three, or five")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("leq 10")
		fallthrough
	case i <= 20:
		fmt.Println("leq 20")
	default:
		fmt.Println("geq 20")
	}

	// Type Switch
	var x interface{} = [3]int{}
	switch x.(type) {
	case int:
		fmt.Println("int")
		break
		// fmt.Println("don't print this")
	case string:
		fmt.Println("string")
	case [2]int:
		fmt.Println("[2]int")
	default:
		fmt.Println("another type")
	}

}

// Summary
// The logical operators are the same as C++ (|| and && and !)
// Go uses short-circuit evaluation
// In switch statements, the break keyword is implied
// Can use fallthrough keyword to get past this which always executes the following case
