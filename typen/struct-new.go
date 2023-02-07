package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	var pv *Vertex

	pv = new(Vertex)
	fmt.Println(pv)

	pv = &Vertex{X:20,Y:21} // impliziert new()
	fmt.Println(pv)
	fmt.Println(*pv)
	fmt.Printf("%T %v %p\n", pv, pv, pv)
	fmt.Printf("%T %v %p\n", *pv, *pv, *pv)
}
