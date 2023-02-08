package main

import (
	"fmt"
	"time"
)

func fibonacci(c chan int, quit chan bool) {
	a, b := 0, 1
	for {
		select {
		case c <- a:
			a, b = b, a+b
		case _, ok := <-quit:
			if !ok {
				return
			}
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan bool)
	go fibonacci(c, quit)

	for i := 0; i < 20; i++ {
		fmt.Printf("%v ", <-c)
	}
	close(quit)

	fmt.Printf("...\n")
	time.Sleep(3 * time.Second) // lasse die Goroutine sich beenden
}
