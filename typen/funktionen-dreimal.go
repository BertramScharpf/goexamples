package main

import "fmt"

func machsDreimal(fn func()) {
	fn("Hans")
	fn("Fred")
	fn("Reto")
}

func main() {
	f := func(n string) {
		fmt.Printf("Hello, %v!", n)
	}
	machsDreimal(f)
}
