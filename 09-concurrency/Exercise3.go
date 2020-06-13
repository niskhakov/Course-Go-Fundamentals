package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	wg.Add(200)
	for i := 0; i < 200; i++ {
		go func() {
			var myCounter = counter
			runtime.Gosched()
			myCounter++
			counter = myCounter
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
