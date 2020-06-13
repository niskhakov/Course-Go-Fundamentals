package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(200)
	for i := 0; i < 200; i++ {
		go func() {
			mutex.Lock()
			var myCounter = counter
			// runtime.Gosched()
			myCounter++
			counter = myCounter
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
