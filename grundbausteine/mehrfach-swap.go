package main

import "fmt"

func main() {
	a, b := 33, 444
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)
}
