package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	ch := make(chan int)

	const gorouts = 10

	var wg sync.WaitGroup
	wg.Add(gorouts)
	for i := 0; i < gorouts; i++ {

		go func(counter int) {
			v := counter
			for i := 0; i < 10; i++ {
				ch <- (10*v + i)
			}
			wg.Done()
		}(i)
	}

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
	fmt.Println("grtns:", runtime.NumGoroutine())
	wg.Wait()
	close(ch)
	fmt.Println("grtns:", runtime.NumGoroutine())
	fmt.Println("Exit")
}
