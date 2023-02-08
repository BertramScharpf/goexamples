package main

import "fmt"

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nichts>")
		return // Ausgabe von t.S wäre tödlich
	}
	fmt.Println(t.S)
}

type I interface {
	M()
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
