package main

import (
	"fmt"
)

func addierer( n int) func (int) int {
	fn := func( x int) int {
		return x + n
	}
	return fn
}

func main() {
	a := addierer(7)
	r := a(5)
	fmt.Println(r)
	fmt.Println(a(8))
	fmt.Println(a(9))
	fmt.Println(a(10))
}

// vim: set noet ts=4 sw=4 sts=4 :
