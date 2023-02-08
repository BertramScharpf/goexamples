package main

import (
	"fmt"
	"time"
)

func sleep_and_signal(n int, c chan bool) {
	time.Sleep(time.Duration(n) * time.Second)
	c <- true
}

func main() {
	fmt.Println("Start")

	c := make(chan bool)
	go sleep_and_signal(3, c)

	select {
	case <-c:
		fmt.Println("Meldung!")
	}

	fmt.Println("Fertig")
}
