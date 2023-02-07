package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}

	var abs func(Vertex) float64 = Vertex.Abs
	fmt.Println(abs(v))

	var a func() float64 = v.Abs
	fmt.Println(a())
}
