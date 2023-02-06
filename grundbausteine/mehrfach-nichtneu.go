package main

import "fmt"

func main() {
	a, b := 33, 444
	fmt.Println(a, b)
	c, b, a := 55, 666, 777 // b war schon deklariert, aber nicht c
	fmt.Println(c, b, a)
}
