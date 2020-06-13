package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go printSmth("Hello", 1)
	go printSmth("Nail", 2)

	wg.Wait()

}

func printSmth(s string, i int) {
	fmt.Println(s, i)
	wg.Done()
}
