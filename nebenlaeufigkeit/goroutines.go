package main

import (
	"fmt"
	"time"
)

func say(s string, pre int) {
	time.Sleep(time.Duration(pre) * time.Millisecond)
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i,s)
	}
}

func main() {
	go say("world", 1000)
	say("hello", 0)
}
