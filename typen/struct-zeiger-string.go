package main

import "fmt"

type Vertex struct {
	X, Y int
}

type Mark struct {
	V     Vertex
	Title string
}

func dontsetit(m Mark) {
	m.V.X, m.V.Y = 33, 44
	fmt.Printf("%v %T\n", &m.Title, &m.Title)
	m.Title = "bar"
	fmt.Printf("%v %T\n", &m.Title, &m.Title)
}

func setit(m *Mark) {
	m.V.X, m.V.Y = 55, 66
	fmt.Printf("%v %T\n", &m.Title, &m.Title)
	m.Title = "baz"
	fmt.Printf("%v %T\n", &m.Title, &m.Title)
}

func main() {
	var m = Mark{Vertex:Vertex{1,2},Title:"foo"}

	fmt.Println(m, m.V)
	dontsetit(m)
	fmt.Println(m, m.V)
	setit(&m)
	fmt.Println(m, m.V)
}
