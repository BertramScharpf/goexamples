package main

import (
	"fmt"
)

var t int = 99

func after() {
	fmt.Println(t)
}

func main() {
	fmt.Println(t)  // Vorsicht. Das ist erlaubt und meint das globale t.
	var t int = 77
	fmt.Println(t)
	after()
}
