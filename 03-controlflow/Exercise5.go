package main

import "fmt"

func main() {
	from := 10
	to := 100
	for i := from; i <= to; i++ {
		fmt.Println(i % 4)
	}
}
