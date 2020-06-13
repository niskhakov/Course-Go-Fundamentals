package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Pink",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "Violet",
		},
		luxury: true,
	}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.color, t1.doors, t1.fourWheel)
	fmt.Println(s1.color, s1.doors, s1.luxury)
}
