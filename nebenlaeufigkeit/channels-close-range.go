package main

import "fmt"

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
	close(c)
}

func main() {
	c := make(chan int)
	go collatz(27, c)

	for v := range c {
		fmt.Println(v)
	}
}
