package main

import "fmt"

func machsDreimal(fn func(string)) {
	fn("Hans")
	fn("Fred")
	fn("Reto")
}

func main() {
	gruss := "Hallo"
	f := func(n string) {
		fmt.Printf("%v, %v!\n", gruss, n)
	}
	machsDreimal(f)
}
