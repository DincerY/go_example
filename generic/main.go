package main

import (
	"example/generic/stack"
	"fmt"
)

func sum[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {

	myStack := stack.CustomStack[int64]{}
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)

	ints := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	floats := map[string]float64{
		"one":   1.2,
		"two":   2.8,
		"three": 3.0,
	}

	fmt.Println(sum(ints))
	fmt.Println(sum(floats))
}
