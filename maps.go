package main

import (
	"fmt"
)

func main() {
	// provincePopulations := make(map[string]int)
	provincePopulations := map[string]int{
		"Ontario":          14570000,
		"Quebec":           8485000,
		"British Columbia": 5071000,
	}
	// the return order of a map is not guaranteed!
	fmt.Println(provincePopulations)
	provincePopulations["Alberta"] = 4371000
	fmt.Println(provincePopulations)
	delete(provincePopulations, "Alberta")
	fmt.Println(provincePopulations)
	fmt.Println(provincePopulations["Alberta"])

	pop, ok := provincePopulations["Ont"]
	fmt.Println(pop, ok) // ok tells us if the key is present in the map
	fmt.Println(len(provincePopulations))

	pp := provincePopulations
	delete(pp, "British Columbia")
	fmt.Println(pp)
	fmt.Println(provincePopulations)
}

// Summary
// Maps can be created via literals or the make function
// KVPs
// Multiple assignments to the same map refer to the same underlying data
