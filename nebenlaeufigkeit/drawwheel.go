package main

import (
	"fmt"
	"time"
)

func DrawWheel(s string) (chan bool) {
	quit := make(chan bool)
	go func() {
		for {
			for _, u := range []rune(s) {
				fmt.Print(string(u))
				select {
				case <-quit:
					fmt.Print("\b \b\n")
					return
				default:
					time.Sleep(150 * time.Millisecond)
				}
				fmt.Print("\b")
			}
		}
	}()
	return quit
}

func main() {
	var q chan bool

	q = DrawWheel(`|/-\`)
	time.Sleep(5 * time.Second)
	close(q)
	time.Sleep(200 * time.Millisecond)

	q = DrawWheel(".oO°*")
	time.Sleep(5 * time.Second)
	close(q)
	time.Sleep(200 * time.Millisecond)
}
