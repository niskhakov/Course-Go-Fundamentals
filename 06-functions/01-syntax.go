package main

import "fmt"

func main() {
	foo()
	bar("Nail")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x, y)

}

func foo() {
	fmt.Println("hello form foo")
}

// everything in Go is PASSED by VALUE
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a, b
}
