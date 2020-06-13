package main

import "fmt"

func main() {
	fmt.Print(foo(2, 3, 4, 5, 6, 7, 8))
}

func foo(x ...int) int {
	fmt.Print(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total:", sum)
	}

	fmt.Println("the total is", sum)
	return sum
}
