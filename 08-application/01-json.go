package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type personObj []struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bss := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bss)

	var peopleObjs []person
	err = json.Unmarshal(bss, &peopleObjs)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", peopleObjs)

	for i, v := range peopleObjs {
		fmt.Println("\nPerson number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
