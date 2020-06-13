package main

import "fmt"

func main() {

	f := func(x int) {
		fmt.Println("My first func expression", x)
	}

	f(322)
}
