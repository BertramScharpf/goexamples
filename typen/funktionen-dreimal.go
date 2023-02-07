package main

import "fmt"

func machsDreimal(fn func()) {
	fn()
	fn()
	fn()
}

func main() {
	f := func() {
		fmt.Println("Hello, here is an anonymous function!")
	}
	machsDreimal(f)
}
