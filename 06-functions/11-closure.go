package main

import "fmt"

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	{
		var y int
		y = 1
		fmt.Println(y)
	}

	// fmt.Println(y) - error

	z := incrementor()
	fmt.Println(z())
	fmt.Println(z())
	fmt.Println(z())

	y := incrementor()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
}
