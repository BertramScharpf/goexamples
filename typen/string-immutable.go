package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "hello"

	// s[1] = "a"     // Fehler: Strings sind "immutable"
	fmt.Println(s)
	t := strings.Replace(s, "e", "a", -1)
	fmt.Println(t)
}
