package main

import "fmt"

func main() {
	course := struct {
		name      string
		attendees int
		lectors   []string
	}{
		name:      "Blockchain 101",
		attendees: 7,
		lectors:   []string{"Lector1", "Lector2"},
	}

	fmt.Println(course)
}
