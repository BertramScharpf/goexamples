package main

import "fmt"

func WordCount(s []string) map[string]int {
	m := make(map[string]int)
	for _, f := range s {
		m[f]++
	}
	return m
}

func main() {
	m := WordCount([]string{"foo", "bar", "foo", "baz", "foo", "bar"})
	fmt.Println(m)
}
