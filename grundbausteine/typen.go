package main

import "fmt"

func main() {
	// v := 42
	// v := 0x7fffffffffffffff
	// v := 0x8000000000000000  // error
	// v := 1.0
	// v := 'ä'
	// v := "xyz"
	v := '€'
	fmt.Printf("v=%v is of type %T\n", v, v)
}
