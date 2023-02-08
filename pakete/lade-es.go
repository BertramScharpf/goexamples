package main

import (
	"fmt"
	"mypkg"
)

func main() {
	var d mypkg.Dummy
	d.PubVal = "hello"
	fmt.Println( d.PubVal)
	mypkg.PubFunc()
}
