package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		// n := i
		go func() {
			time.Sleep(2 * time.Millisecond)
			fmt.Println("Hallo Nr.", i)
		}()
	}
	time.Sleep(5 * time.Second)
}
