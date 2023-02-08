package main

import (
	"fmt"
	"time"
)

func say(s string, delay int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go func() {
		say("world", 200)
	}()
	say("hello,", 100)
	time.Sleep(5 * time.Second)
}
