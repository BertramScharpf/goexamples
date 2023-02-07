package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	tmp := make([]int, len(arr)+1)
	copy(tmp, arr)
	fmt.Println(tmp)
}
