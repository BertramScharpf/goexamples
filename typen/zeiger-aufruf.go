package main

import "fmt"

func modint( i int) {
	i = 53
}

func modpint( p *int) {
	*p = 64
}

func main() {
	i := 42

	modint(   i)
	fmt.Printf("i = %v\n", i)
	modpint( &i)
	fmt.Printf("i = %v\n", i)
}

