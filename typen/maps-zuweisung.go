package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The Value:", m["Answer"])
	// m["Answer"] = 0
	for k := range m {
		fmt.Println(k)
	}

	v, ok := m["Answer"]
	if ok {
		fmt.Println("The value:", v)
	} else {
		fmt.Println("Kein Wert für \"Answer\".")
	}
}
