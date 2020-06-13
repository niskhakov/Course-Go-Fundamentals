package main

import "fmt"

func main() {
	w := wtfer()

	w("Hello")
}

func wtfer() func(string) {
	return func(s string) {
		fmt.Println(s, "wtf!")
	}
}
