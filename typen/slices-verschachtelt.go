package main

import "fmt"

func Pic(dx, dy uint) [][]uint8 {

	p := make([][]uint8, dy)
	fmt.Println(p)
	for i := range p {
		p[i] = make([]uint8, dx)
	}
	fmt.Println(p)
	fmt.Println("----")

	for y, row := range p {
		for x := range row {
			row[x] = uint8(x * y)
		}
	}

	return p
}

func main() {
	p := Pic(8, 9)
	for _, row := range p {
		fmt.Println(row)
	}
}
