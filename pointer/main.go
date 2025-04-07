package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1
	p.Y = 2

	var a *int
	value := 10

	a = &value

	fmt.Println(*a)
	*a = 21
	fmt.Println(value)

	/*if a == nil {
		fmt.Printf("pointer is nil")
	} else {
		fmt.Printf("pointer is not nil")
	}*/

}
