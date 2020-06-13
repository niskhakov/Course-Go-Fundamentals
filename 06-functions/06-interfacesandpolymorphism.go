package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func bar(h human) {
	// asserting
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)

}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: true,
	}
	sa1.speak()
	sa2.speak()

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
