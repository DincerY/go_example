package main

import "fmt"

func main() {

	map1 := map[string]int{"a": 1, "b": 2}
	map2 := make(map[string]int, 5)

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 5)

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(slice1)
	fmt.Println(slice2)

}
