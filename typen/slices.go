package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%v (%T)\n", primes, primes)

	var s []int = primes[1:4]
	fmt.Printf("%v (%T)\n", s, s)
	fmt.Println(s)
}
