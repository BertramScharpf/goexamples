package main

import "fmt"

func main() {
	var v int = 0x7fffffffffffffff
	var w float64 = float64(v * v)
	fmt.Println(w)
	var q float64 = float64(v) * float64(v)
	fmt.Println(q)
}
