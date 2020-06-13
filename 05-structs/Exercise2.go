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

	persons := map[string]person{p1.last: p1, p2.last: p2}

	for k, v := range persons {
		fmt.Println(v.first, k)
		for _, v := range p1.icecreamFlavours {
			fmt.Print(v, " ")
		}
		fmt.Println("")
	}
}
