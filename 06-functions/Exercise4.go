package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Nail",
		last:  "Iskhakov",
		age:   22,
	}

	p.speak()
}
