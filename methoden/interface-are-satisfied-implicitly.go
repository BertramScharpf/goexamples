package main

import "fmt"

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type I interface {
	M()
}

func main() {
	var t T = T{"hello"}
	var i I = &t
	i.M()
}
