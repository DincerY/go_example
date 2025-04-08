package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	go Walk(t1, ch)
	res := true
	test(t2, ch, &res)
	return res
}

func test(t *tree.Tree, ch chan int, res *bool) {
	if t == nil {
		return
	}
	test(t.Left, ch, res)
	if t.Value != <-ch {
		*res = false
	}
	test(t.Right, ch, res)
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func main() {

	fmt.Println(Same(tree.New(2), tree.New(2)))

}
