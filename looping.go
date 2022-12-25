package main

import (
	"fmt"
)

func main() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// There's no while loop, instead can leave off some commands:
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}

	// Instead of using flagging, Go has a built in feature called Labels
Loop:
	for u := 1; u <= 3; u++ {
		for v := 1; v <= 3; v++ {
			fmt.Println(u * v)
			if u * v >= 3 {
				break Loop
			}
		}
	}

	// Iterating through a slice/array (same technique for maps and strings)
	s := []int{1, 2, 3}
	fmt.Println(s)
	for k, v := range s {
		fmt.Println(k, v)
	}

	// Note that you must use all declared variables, so use '_' to let the compiler know you don't need it
	st := "Hello World"
	for _, v := range st {
		fmt.Println(string(v))
	}
}