package main

import "fmt"

func main() {
	name := "Egor"
	admin := "Nail"

	switch {
	case (42 == 32+10):
		fmt.Println("Okay, that's fun", name)
	case 34 < 42:
		fmt.Println("Okay, that's true", admin)
	case 12 > 45:
		fmt.Println("No way")
	}
}
