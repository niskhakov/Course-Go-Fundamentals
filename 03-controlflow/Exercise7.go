package main

import "fmt"

func main() {
	name := "Egor"
	admin := "Nail"
	moderator := "Nils"

	if name == admin {
		fmt.Println("Access authorized")
	} else if name == moderator {
		fmt.Println("Oh, you are moderator, please come here, there is a control unit of nuclear missiles")
	} else {
		fmt.Println("You are not an admin, you have noo power here")
	}
}
