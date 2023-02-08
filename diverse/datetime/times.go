package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Duration = 3*time.Hour +
		4*time.Second +
		5*time.Microsecond
	fmt.Println(t)
	t = 5*time.Microsecond
	fmt.Println(t)
}
