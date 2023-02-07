package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) String() string {
	return fmt.Sprintf("{X=%v Y=%v}", v.X, v.Y)
}

func main() {
	fmt.Println(Vertex{1, 2})
}
