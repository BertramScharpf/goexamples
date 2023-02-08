package main

import "fmt"

func wie_python2() {
	var l []int
	for i := 0; i < 1000; i++ {
		l = append(l, i)
	}

	for i := range l {
		fmt.Printf("%d ", i*i)
	}
	fmt.Printf("\n")
}

func wie_python3() {
	c := make(chan int)

	go func(max int, d chan int) {
		for i := 0; i < max; i++ {
			d <- i
		}
		close(d)
	}(1000, c)

	for i := range c {
		fmt.Printf("%d ", i*i)
	}
	fmt.Printf("\n")
}

func main() {
	wie_python3()
}
