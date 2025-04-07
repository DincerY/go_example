package main

import (
	"fmt"
)

type AnimalSounder interface {
	MakeNoise()
}

type Animal struct {
	Name   string
	Gender string
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
}

func main() {
	cat := &Cat{Animal{Name: "Gizmo", Gender: "girl"}}
	dog := &Dog{Animal{Name: "Pablo", Gender: "boy"}}

	makeAnimalNoise(cat)
	makeAnimalNoise(dog)
}

func (a *Animal) PerformNoise(name, gender, perform string) {
	fmt.Printf("%s is a %s and %s", name, gender, perform)
	fmt.Println("")
}

func (c *Cat) MakeNoise() {
	c.Animal.PerformNoise(c.Name, c.Gender, "meowing")
}

func (d *Dog) MakeNoise() {
	d.Animal.PerformNoise(d.Name, d.Gender, "barking")
}

func makeAnimalNoise(sounder AnimalSounder) {
	sounder.MakeNoise()
}
