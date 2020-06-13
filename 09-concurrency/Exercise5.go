package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	wg.Add(200)
	for i := 0; i < 200; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			// counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
