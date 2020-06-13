package main

import "fmt"

func main() {
	x1 := sliceScope()
	x2 := sliceScope()

	x1(10)
	x1(42)
	fmt.Println(x1(23))
	x2(2)
	x2(5)
	fmt.Println(x2(11))
	x2(7)
	fmt.Println(x2(10))

}

func sliceScope() func(int) []int {
	x := []int{}
	return func(y int) []int {
		x = append(x, y)
		return x
	}
}
