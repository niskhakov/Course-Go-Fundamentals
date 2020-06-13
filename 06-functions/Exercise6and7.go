package main

import "fmt"

func main() {
	printer := func(msg string) {
		fmt.Println(msg)
	}

	printer("Page print 1")
}
