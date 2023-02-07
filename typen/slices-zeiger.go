package main

import "fmt"

func main() {
	names := [4]string{ "John", "Paul", "George", "Ringo", }
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	fmt.Println()

	b[0] = "XXX"
	fmt.Println(b)
	fmt.Println(names)
	fmt.Println(a)
}
