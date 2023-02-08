package main

import (
	"fmt"
	"time"
)

func after(n int) chan bool {
	b := make(chan bool)
	go func(b chan bool) {
		time.Sleep(time.Duration(n) * time.Second)
		close( b)
	}(b)
	return b
}

func timer(n int) chan bool {
	c := make(chan bool)
	go func(c chan bool) {
		for {
			time.Sleep(time.Duration(n) * time.Second)
			c <- true
		}
	}(c)
	return c
}

func main() {
	t := timer(1)
	b := after(5)
	for running := true; running; {
		select {
		case <-t:
			fmt.Print("tick")
		case _, running = <-b:
			fmt.Print("BOOM!")
		default:
			fmt.Print(".")
			time.Sleep(time.Duration(50) * time.Millisecond)
		}
	}
	fmt.Println()
}
