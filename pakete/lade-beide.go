package main

import (
	"mypkg"
)

func main() {
	var d mypkg.Dummy
	var o mypkg.Other
	d.PubVal = "hello"
	o.Title = "foo"
}
