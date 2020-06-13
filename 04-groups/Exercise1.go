package main

import "fmt"

func main() {
	a := [5]int{1, 3, 5, 7, 9}
	// a[0] = 1
	// a[1] = 3
	// a[2] = 5
	// a[3] = 7
	// a[4] = 9

	for i, v := range a {
		fmt.Printf("index: %d -> %d - %T\n", i, v, a)
	}
}
