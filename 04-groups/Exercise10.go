package main

import "fmt"

func main() {
	db := map[string][]string{
		"Nail":      []string{"dumb things", "frontend", "architecture", "art"},
		"Todd":      []string{"course creation", "go programming"},
		"Alexander": []string{"blockchain", "management"},
	}

	db["Grisha"] = []string{"cool", "guy"}

	for k, v := range db {
		fmt.Println(k, v)
	}

	delete(db, "Grisha")

	for k, v := range db {
		fmt.Println(k, v)
	}

}
