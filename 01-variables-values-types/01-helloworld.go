package main

import "fmt"

func main()  {
	fmt.Println("Hello everyone")
	foo()
	fmt.Println("something new")

	for i:=0; i< 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("Im'n fooo")
}

func bar()  {
	fmt.Println("and then we exited")
}
// control flow:
// (1) sequence
// (2) loops
// (3) conditional