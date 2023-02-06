package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Printf("p ist ein %T mit dem Wert %v\n", p)
	fmt.Println(*p)
	*p = 21
	fmt.Printf("i = %v\n", i)

	p = &j
	*p /= 37
	fmt.Printf("j = %v\n", j)
}
