package main

import "fmt"

func main() {
	var a rune = 'A'
	var z rune = 'Z'

	for i := a; i <= z; i++ {
		for k := 0; k < 3; k++ {
			fmt.Printf("%#U \n", i)
		}
	}
}
