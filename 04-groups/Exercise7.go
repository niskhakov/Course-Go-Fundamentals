package main

import "fmt"

func main() {
	s := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooo, James"},
	}

	for _, out := range s {
		for _, inn := range out {
			fmt.Println(inn)
		}
	}
}
