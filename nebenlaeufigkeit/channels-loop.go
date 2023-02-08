package main

import (
	"fmt"
)

func provide(c chan int) {
	n := 1
	for i := 0; i < 10; i++ {
		n *= 2
		c <- n - 1
	}
}

func main() {
	c := make(chan int)
	go provide(c)

	for i := 0; i < 9; i++ { // bis 11 bräche ab
		fmt.Println(<-c)
	}
}
