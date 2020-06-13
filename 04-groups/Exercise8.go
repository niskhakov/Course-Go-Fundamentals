package main

import "fmt"

func main() {
	db := map[string][]string{
		"Nail":      []string{"dumb things", "frontend", "architecture", "art"},
		"Todd":      []string{"course creation", "go programming"},
		"Alexander": []string{"blockchain", "management"},
	}

	for k, v := range db {
		fmt.Println(k, v)
	}

}
