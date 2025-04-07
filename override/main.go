package main

import "fmt"

type Test struct {
	x int32
	y int32
}

func (t Test) String() string {
	return "override string func"
}

func main() {
	newT := new(Test)
	fmt.Println(newT)
	t := Test{1, 2}
	fmt.Println(t)
}
