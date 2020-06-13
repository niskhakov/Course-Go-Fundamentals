package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	for i, v := range a {
		fmt.Printf("index: %d -> %d - %T\n", i, v, a)
	}
}
