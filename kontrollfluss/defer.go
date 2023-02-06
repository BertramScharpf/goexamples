package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("world %v", i)
	}
	fmt.Println("hello")
	time.Sleep(3 * time.Second)
}
