package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v

	(*p).X = 20
	p.X = 20 // gleichwertige Schreibweise

	fmt.Println(v)
}
