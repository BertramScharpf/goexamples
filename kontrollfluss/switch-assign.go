package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{3, 4}
	switch v {
	case Vertex{3, y}:  // Das geht in vielen Sprachen, aber nicht in Go.
		fmt.Printf("Trifft mit y=%v.", y)
	default:
		fmt.Println("Trifft nicht.")
	}
}
