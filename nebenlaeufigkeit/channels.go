package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "foo"
	c <- "bar"
	fmt.Println(<-c)
	c <- "baz"
	fmt.Println(<-c)
	fmt.Println(<-c)
}
