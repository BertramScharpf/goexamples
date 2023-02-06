package main

import "fmt"

func rot(x, y, z string) (string, string, string) {
	return z, x, y
}

func main() {
	a, b, c := rot("world", "!", "hello")
	fmt.Println(a, b, c)
	// fmt.Println("a", "b", "c")
	// fmt.Println("a" + "b" + "c")
}
