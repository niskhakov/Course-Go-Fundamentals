package main

import "fmt"

func main() {
	var favSport string = "swimming"

	switch favSport {
	case "cycling":
		fmt.Println("Okay, cycling is good")
	case "running":
		fmt.Println("Okay, running is awesome")
	case "swimming":
		fmt.Println("No way")
	default:
		fmt.Println("That's also good")
	}
}
