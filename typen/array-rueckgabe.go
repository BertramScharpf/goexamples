package main

import "fmt"

func dummies() [3]string {
	var b [3]string = [...]string{"foo", "bar", "baz"}
	return b
}

func ptrdummies() *[3]string {
	var b [3]string = [...]string{"FOO", "BAR", "BAZ"}
	return &b
}

func main() {
	a := dummies()
	fmt.Printf("%v (%T)\n", a, a)
	fmt.Println(a)

	b := ptrdummies()
	fmt.Printf("%v (%T)\n", b, b)
	fmt.Printf("%v (%T)\n", *b, *b)
	fmt.Println(*b)
}
