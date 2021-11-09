package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLenght float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{sideLenght: 2}
	tr := triangle{base: 3, height: 4}

	printArea(sq)
	printArea(tr)
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLenght, 2)
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
