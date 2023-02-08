//
//  onlyone.go  --  Nur eine Funktion überschreiben
//

package main

import (
	"fmt"
)

// Some
type Some struct {
	M int
}

func (_ Some) foo() { fmt.Println("Some's foo") }
func (_ Some) bar() { fmt.Println("Some's bar") }
func (_ Some) baz() { fmt.Println("Some's baz") }

// Many
type Many interface {
	foo()
	bar()
	baz()
}

// More
type More struct {
	S      Some // !!!
	Status int
}

func (r *More) foo() { fmt.Println("More's foo"); r.Status = 200 }

// main
func main() {
	var s Some

	var p *Some
	p = &s
	p.foo()

	var m Many
	m = &s
	m.foo()

	var r More = More{s, 0}
	m = &r
	m.foo()
	m.bar()
	fmt.Println(r.Status)
}

// vim: set noet ts=4 sw=4 sts=4 :
