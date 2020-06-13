package main

import "fmt"

type person struct {
	first            string
	last             string
	icecreamFlavours []string
}

func main() {
	p1 := person{
		first:            "Nail",
		last:             "Iskhakov",
		icecreamFlavours: []string{"truffle", "chocolate"},
	}

	p2 := person{
		first:            "Andrei",
		last:             "Nagoe",
		icecreamFlavours: []string{"nut", "cinnabon"},
	}

	fmt.Println(p1.first, p1.last)
	for _, v := range p1.icecreamFlavours {
		fmt.Print(v, " ")
	}
	fmt.Println("")
	fmt.Println(p2.first, p2.last)
	for _, v := range p2.icecreamFlavours {
		fmt.Print(v, " ")
	}
}
