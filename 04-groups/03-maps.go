package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v, ok)

	if v, ok := m["James"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v, ok)
	}

	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{1, 2, 3, 5, 6}
	for i, v := range xi {
		fmt.Println(i, v)
	}

	delete(m, "todd")
	delete(m, "nail")
	for k, v := range m {
		fmt.Println(k, v)
	}

	if v, ok := m["James"]; ok {
		fmt.Println("value", v)
		delete(m, "James")
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
