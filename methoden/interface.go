package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

func main() {
	v := Vertex{3, 4}

	var a Abser
	a = &v
	fmt.Println(a.Abs())

	// Nur der Zeiger ist ein Abser
	// a = v
	// fmt.Println(a.Abs())
}
