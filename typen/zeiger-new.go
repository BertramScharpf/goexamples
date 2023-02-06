package main

import "fmt"

func main() {
	var p *int
	p = new(int)
	fmt.Printf("p: %v an %p (%T)\n", *p, p, p)
}
