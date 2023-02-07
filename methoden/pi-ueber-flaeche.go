package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

const (
	Num = 0x1000000
)

func main() {
	inside := 0
	for i := 0; i < Num; i++ {
		v := Vertex{rand.Float64(), rand.Float64()}
		if v.Abs() < 1.0 {
			inside++
		}
	}
	var res float64 = 4.0
	res *= float64(inside)
	res /= Num
	fmt.Println(res)
}
