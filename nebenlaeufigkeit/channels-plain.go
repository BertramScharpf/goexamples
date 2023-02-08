package main

import (
	"fmt"
)

func provide(c chan string) {
	fmt.Println("SDtarte Lieferung.")
	c <- "foo"
	fmt.Println("foo geliefert")
	c <- "bar"
	fmt.Println("bar geliefert")
	c <- "baz"
	fmt.Println("baz geliefert")
}

func main() {
	c := make(chan string)
	go provide(c)

	fmt.Println("Hole erstes")
	fmt.Println(<-c)
	fmt.Println("Hole zweites")
	fmt.Println(<-c)
	fmt.Println("Hole drittes")
	fmt.Println(<-c)
	fmt.Println("Fertiggeholt.")
}
