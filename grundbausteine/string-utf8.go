package main

import "fmt"

func main() {
	s := "Hallöchen"
	r := []rune(s)

	println(r)
	fmt.Println(r)
	fmt.Println(string(r[4:6]))
	fmt.Println(string(r[4:5]))
	fmt.Println(string(r[:4]))
	fmt.Println(string(r[5:]))
}
