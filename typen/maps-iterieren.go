package main

import "fmt"

var m = map[string]int{
	"foo": 33,
	"bar": 44,
	"baz": 55,
}

func main() {
	for k := range m {
		fmt.Println(k)
	}
}
