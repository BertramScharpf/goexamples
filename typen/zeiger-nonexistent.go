// Anzeigen lassen, daß die Variable in den Heap verlagert wird:
// go build -gcflags -m typen/zeiger-nonexistent.go

package main

import "fmt"

func gimmeptr() *int {
	var i int = 99
	return &i
}

func main() {
	p := gimmeptr()
	fmt.Printf("*p = %v\n", *p)
}
