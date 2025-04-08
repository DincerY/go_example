package main

import "fmt"

func main() {
	test(1, 2, 3, 4, 5, 6)
	test(1, 2)
	test(1, 2, 3)
	test(1, 2, 3, 4)

	variables := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	test(variables...)
}

//nums parametresi slice olarak fonksiyona iletilir
func test(nums ...int) {
	fmt.Println(len(nums))
}
