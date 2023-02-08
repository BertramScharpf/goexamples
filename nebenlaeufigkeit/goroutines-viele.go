package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(n int) {
			time.Sleep(2 * time.Millisecond)
			fmt.Println("Hallo Nr.", n)
		}(i)
	}
	time.Sleep(5 * time.Second)
}
