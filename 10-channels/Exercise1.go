package main

import (
	"fmt"
)

func mainAnon() {
	c := make(chan int)

	go func() { c <- 42 }()

	fmt.Println(<-c)
}

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
