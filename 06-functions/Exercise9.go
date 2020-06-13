package main

import "fmt"

func main() {
	post("Hello", func(s string) bool {
		fmt.Println(s)
		return true
	})
}

func post(s string, next func(s string) bool) {
	fmt.Println(s)
	si := fmt.Sprint(s, "Nail!")
	fmt.Println(next(si))
}
