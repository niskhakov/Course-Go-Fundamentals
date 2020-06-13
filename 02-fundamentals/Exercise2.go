package main

import "fmt"

func main() {
	x1 := 42 == 42
	x2 := 42 <= 43
	x3 := 42 >= 43
	x4 := 42 < 42
	x5 := 42 > 41

	fmt.Println(x1, x2, x3, x4, x5)
}
