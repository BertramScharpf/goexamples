package main

import "fmt"

func swap( p, q *int) {
	*p, *q = *q, *p
}

func main() {
	i, j := 42, 33

	p, q := &i, &j
	fmt.Printf("*p, *q = %v, %v\n", *p, *q)
	p, q = &j, &i
	fmt.Printf("*p, *q = %v, %v\n", *p, *q)
	p, q = q, p
	fmt.Printf("*p, *q = %v, %v\n", *p, *q)
	fmt.Printf("i,  j  = %v, %v\n", i,  j)
	fmt.Printf("-----------------\n")
	swap(&i,&j)
	fmt.Printf("i,  j  = %v, %v\n", i,  j)
}

