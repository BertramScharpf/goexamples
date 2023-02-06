package main

import "fmt"

func force_panic() {
	var a []int = make([]int,0)
	fmt.Println(a[2])
}

func main() {
	defer fmt.Println("Muß noch aufräumen...")
	force_panic()
	fmt.Println("Tschüß!")
}
