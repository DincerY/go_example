package main

import (
	"fmt"
	"math"
)

//A method is a function with a special receiver argument.

type Vertex struct {
	x int
	y int
}

func (vertex Vertex) Abs() int {
	return int(math.Sqrt(float64(vertex.x) + float64(vertex.y)))
}

func (vertex *Vertex) UpdateX(x int) {
	vertex.x = x
}

func (vertex *Vertex) UpdateY(y int) {
	vertex.y = y
}

func main() {
	vertex := Vertex{x: 1, y: 2}

	/*vertex.UpdateX(10)
	vertex.UpdateY(10)*/
	fmt.Println(vertex.Abs())

}
