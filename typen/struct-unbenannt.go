package main

import "fmt"

func main() {
	s := struct {
		anzahl int
		text   string
	}{3, "hi"}
	fmt.Println(s)
}
