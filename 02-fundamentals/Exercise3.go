package main

import "fmt"

const (
	a     = 12
	b int = 14
)

func main() {
	fmt.Println(a, b)
	fmt.Printf("%T, %T\n", a, b)
}
