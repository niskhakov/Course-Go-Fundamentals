package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type UserByAge []user

func (a UserByAge) Len() int           { return len(a) }
func (a UserByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a UserByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type UserByLast []user

func (a UserByLast) Len() int           { return len(a) }
func (a UserByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a UserByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func print(u user) {
	fmt.Printf("\t%s %s %d\n", u.First, u.Last, u.Age)
	for i, v := range u.Sayings {
		fmt.Println("\t\t", i, "->", v)
	}
}

func printUsers(u []user) {
	for i, v := range u {
		fmt.Println("Person:", i)
		print(v)
	}
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	users := []user{u1, u2}
	sort.Strings(u1.Sayings)
	sort.Strings(u2.Sayings)

	fmt.Println("Sort by Age:")
	sort.Sort(UserByAge(users))
	printUsers(users)

	fmt.Println("Sort by Last:")
	sort.Sort(UserByLast(users))
	printUsers(users)

}
