package main

import "fmt"

func main() {
	c := make(chan int, 1)
	// c := make(chan int, 2)

	c <- 42
	// program blocks, cause there is no room for 2nd int, so it waits until first int been read
	c <- 43

	fmt.Println(<-c)
	// fmt.Println(<-c )
}
