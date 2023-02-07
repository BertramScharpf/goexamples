package main

import "fmt"

func main() {
	var i int8 = -125
	fmt.Println(uint8(i*i))
	fmt.Println(i)
	fmt.Println(uint8(i))
	fmt.Println(uint8(i)*uint8(i))
}
