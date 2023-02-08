package main

import (
	"fmt"
)

func collatz(n int, c chan int) {
	for {
		c <- n
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	c := make(chan int)
	go collatz(17, c)

	for i := 0; i < 9; i++ {
		fmt.Println(<-c)
	}
}
