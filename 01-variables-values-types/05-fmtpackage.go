package main

import "fmt"

var y = 42

func main()  {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	s := fmt.Sprintf("%#x\n%b\n%x\n", y,y,y)
	fmt.Print(s)

	fmt.Printf("%v", y)
}
