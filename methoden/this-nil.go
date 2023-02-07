package main

import "fmt"

type dummy struct{}

func (this *dummy) doit() {
	if this == nil {
		println("ja")
	} else {
		println("nein")
	}
}

func main() {
	var d *dummy
	fmt.Printf("%v (%T)\n", d, d)
	d.doit()
}
