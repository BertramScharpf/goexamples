package main

import "fmt"

func main() {
	var a [3]interface{}
	a[0] = "Hello"
	a[1] = 3.14159
	a[2] = false
	fmt.Println(a)
}
