package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hello, here is an anonymous function!")
	}
	fmt.Printf("%v (%T)\n", f, f)
	f()
}
