package main

import (
	"fmt"
	"time"
)

func objdummies() [3]string {
	var b [3]string = [...]string{"foo", "bar", "baz"}
	go func() {
		time.Sleep(2 * time.Second)
		b[1] = "BAR"
	}()
	return b
}

func ptrdummies() *[3]string {
	var b [3]string = [...]string{"FOO", "BAR", "BAZ"}
	go func() {
		time.Sleep(2 * time.Second)
		b[1] = "Bar"
	}()
	return &b
}

func main() {
	a := objdummies()
	fmt.Printf("%v (%T)\n", a, a)
	fmt.Println(a)

	b := ptrdummies()
	fmt.Printf("%v (%T)\n", b, b)
	fmt.Printf("%v (%T)\n", *b, *b)
	fmt.Println(*b)

	time.Sleep(5 * time.Second)
	fmt.Printf("%v (%T)\n", a, a) // hat sich nicht verändert
	fmt.Printf("%v (%T)\n", b, b) // hat sich verändert
}
