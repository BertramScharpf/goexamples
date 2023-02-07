package main

import "fmt"

func slcdummies() []string {
	var b [3]string = [...]string{"FOO", "BAR", "BAZ"}
	return b[:1]
}

func main() {
	b := slcdummies()
	fmt.Printf("%v (%T)\n", b, b)
	fmt.Printf("len=%d cap=%d\n", len(b), cap(b))
}
