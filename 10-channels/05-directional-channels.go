package main

import "fmt"

func main() {
	// send-only channel
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	fmt.Printf("%T\n", c)

	// receive-only channel
	r := make(<-chan int, 2)

	fmt.Printf("%T\n", r)

	// example
	ch := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	fmt.Println("--------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// general to specific assigns does work
	cr = ch
	cs = ch

}
