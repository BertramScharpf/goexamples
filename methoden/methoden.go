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

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	w := Vertex{20, 21}
	fmt.Println(v.Abs())
	fmt.Println(w.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

    p := &w
	fmt.Println(p.Abs())
	p.Scale(10)
	fmt.Println(p.Abs())
}
