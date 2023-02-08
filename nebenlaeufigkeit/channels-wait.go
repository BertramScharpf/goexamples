package main

import (
	"fmt"
	"time"
)

func provide(c chan string) {
	fmt.Println("sending foo.")
	c <- "foo"
	fmt.Println("sending bar.")
	c <- "bar"
	fmt.Println("sending baz.")
	c <- "baz"
}

func main() {
	c := make(chan string)
	go provide(c)

	time.Sleep(1 * time.Second)
	fmt.Println(<-c)
	time.Sleep(1 * time.Second)
	fmt.Println(<-c)
	time.Sleep(1 * time.Second)
	fmt.Println(<-c)
}
