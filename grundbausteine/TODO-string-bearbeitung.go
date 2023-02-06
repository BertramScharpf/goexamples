package main

import "fmt"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func h(s string) string {
	return s
}

func main() {
	s := h("hello")
	r := []rune(s)
	fmt.Println(r)
	fmt.Println(Reverse("Reliefpfeiler"))
}
