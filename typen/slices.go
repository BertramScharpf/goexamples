package main

import "fmt"

func main() {
	primes := [...]int{2, 3, 5, 7, 11, 13}
	// primes:         0__1__2__3__4___5_
	// s:					 0__1__2_
	fmt.Printf("%v (%T)\n", primes, primes)

	var s []int = primes[2:5]
	fmt.Printf("%v (%T)\n", s, s)
	fmt.Println(s)

	primes[1] = 37
	fmt.Printf("%v (%T)\n", primes, primes)
	fmt.Printf("%v (%T)\n", s, s)

	s[2] = 73
	fmt.Printf("%v (%T)\n", primes, primes)
	fmt.Printf("%v (%T)\n", s, s)
}
