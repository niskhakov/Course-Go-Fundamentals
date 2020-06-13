package main

import "fmt"

func main() {
	name := "Egor"
	admin := "Nail"

	if name == admin {
		fmt.Println("Access authorized")
	} else {
		fmt.Println("You are not an admin, you have noo power here")
	}
}
