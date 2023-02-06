package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 30, 4
	fmt.Println(float64(x)/float64(y))
	fmt.Println(float64(x/y))
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func main_float() {
	var x, y float64 = 30, 4
	fmt.Println(x/y)
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
