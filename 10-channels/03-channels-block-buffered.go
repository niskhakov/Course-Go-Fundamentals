package main

import "fmt"

func main() {
	// make 1 int sit in the channel without block
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
