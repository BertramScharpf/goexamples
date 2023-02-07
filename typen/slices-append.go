package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4) // Kapazität wächst mehr
	printSlice(s)

	s = append(s, 5, 6, 7, 8, 9) // Kapazität wächst mehr
	printSlice(s)
}
