package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) changeMe(newName string) {
	(*p).name = newName
}

func main() {
	val := 12
	fmt.Println(&val, val)

	p := person{
		name: "Nail Iskhakov",
		age:  22,
	}
	fmt.Println(p)
	p.changeMe("Nail ???")
	fmt.Println(p)
}
