package main

import (
	"fmt"
)

type TestError struct {
	x int
	y int
}

func (te TestError) Error() string {
	return "bu bir denemedir"
}

func testSqrt() (int, error) {
	return 10, TestError{}
}

// Burada interface benzetmesi yapıyoruz
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	fmt.Println(testSqrt())

}
