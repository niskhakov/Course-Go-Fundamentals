package main

import "fmt"

func main() {
	foo()
}

func foo() {

	msg("Start")
	defer msg("End")
	msg("After defer")

}

func msg(s string) {
	fmt.Println(s)
}
