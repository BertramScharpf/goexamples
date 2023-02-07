package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Mark struct {
	Vertex
	Title string
}

func main() {
	m := Mark{Vertex{1,2}, "A"}
	fmt.Println(m)
	fmt.Println(m.X)

	m = Mark{Vertex:Vertex{X:2,Y:3},Title:"A"}
	fmt.Println(m)
	fmt.Println(m.X)

	m = Mark{Title:"A"}
	fmt.Println(m)
	fmt.Println(m.X)
	fmt.Println(m.Vertex.X)
}
