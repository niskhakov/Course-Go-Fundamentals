package main

import "fmt"

func main() {
	// x := type{values} // composite literal
	x := []int{4, 5, 6, 7, 8, 42}
	fmt.Println(x, len(x), cap(x))
	fmt.Printf("%T\n", x)

	fmt.Println(x[0], x[1], x[2], x[3], x[4], x[5])

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println(x[0:3])
	fmt.Println(x[:])
	fmt.Println(x[4:])

	x = append(x, 1, 2, 3)
	fmt.Println(x)

	y := []int{234, 456, 789, 987}
	x = append(x, y...)
	fmt.Println(x)

	z := append(x[4:], x[:2]...)
	fmt.Println(z)

	x = make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[9] = 9999
	x[0] = 11

	x = append(x, 42)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	jb := []string{"James", "Bond", "Chocoalate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "MoneyPenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
