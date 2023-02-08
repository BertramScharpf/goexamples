package main

import (
	"fmt"
	"time"
)

func collatz(n int, c chan int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panik innen abgefangen.")
        }
	}()
	c <- n
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		c <- n
	}
	time.Sleep(1 * time.Second)
	panic("stop innen")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panik außen abgefangen.")
        }
	}()
	c := make(chan int)
	go collatz(5, c)

	for i := 0; i < 6; i++ {
		fmt.Println(<-c)
	}
	time.Sleep(3 * time.Second)
	panic("stop außen")
}
