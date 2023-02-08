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

	// for done := false; !done; {
	Loop: for {
		select {
		case <-c:
			fmt.Println("Meldung!") // Wird hier übergangen.
			// done = true
			break Loop
		default:
			fmt.Println("Nichts ist passiert.")
			time.Sleep(500 * time.Millisecond)
		}
	}

	fmt.Println("Fertig")
}
