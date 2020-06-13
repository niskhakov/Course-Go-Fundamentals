package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Pass by value: ")
	foo(x)
	fmt.Println(x)

	fmt.Println("Pass pointer:")

	fmt.Println("&x before", &x)
	fmt.Println("x before", x)
	foo1(&x)
	fmt.Println("&x after", &x)
	fmt.Println("x after", x)
}

// pass by value
func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func foo1(y *int) {
	fmt.Println(y)
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)
}
