/*
 *  scanln.go  --  Von der Standardeingabe lesen
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter text: ")
	var t string
	fmt.Scanf("%s", &t)
	fmt.Println(t)

	fmt.Print("Enter number: ")
	var i int
	fmt.Scanf("%d", &i)    // Vorsicht: Rest bleibt in der Warteschlange.
	fmt.Println(i)
}

// vim: set noet ts=4 sw=4 sts=4 :
