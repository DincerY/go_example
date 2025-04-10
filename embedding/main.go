package main

import (
	"fmt"
)

type Base struct {
	baseName string
	baseVal  int
}

func (base Base) GetBaseName() string {
	return base.baseName
}

type Derived struct {
	Base
}

type Baser interface {
	GetBaseName() string
}

func main() {
	derived := Derived{}
	derived.Base.baseName = "Test"

	fmt.Println(derived.GetBaseName())

	var b Baser = Derived{
		Base: Base{
			baseName: "Test 2",
		},
	}
	fmt.Println(b.GetBaseName())
}
