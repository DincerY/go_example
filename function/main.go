package main

import "fmt"

func compute(fn func(x, y int) int) int {
	return fn(2, 3)
}

func main() {

	hypot := func(x, y int) int {
		return x + y
	}

	fmt.Println(hypot(1, 2))
	fmt.Println(compute(hypot))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}

func fibonacci() func() int {
	first := 0
	second := 1
	return func() int {
		val := first + second
		first = second
		second = val
		return first
	}
}
