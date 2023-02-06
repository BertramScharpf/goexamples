/*
 *  ueberlauf.go  --  Integer-Überlauf
 */

package main

import (
	"fmt"
)

func main() {
	var i int8 = 125
	i++
	fmt.Printf("%v\n", i)
	i++
	fmt.Printf("%v\n", i)
	i++
	fmt.Printf("%v\n", i)
	i++
	fmt.Printf("%v\n", i)
}

// vim: set noet ts=4 sw=4 sts=4 :
