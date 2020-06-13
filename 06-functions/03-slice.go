package main

import "fmt"

func main() {
	xs := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := sum(xs...)
	fmt.Print(x)
}

func sum(x ...int) int {
	fmt.Print(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}

	fmt.Println("the total is", sum)
	return sum
}
