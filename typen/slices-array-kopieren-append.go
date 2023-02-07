package main

import "fmt"

const N = 1024

func main() {
	arr := make([]uint,N)
	for i := range arr {  // i ist int, nicht uint
		arr[i] = uint(i)*uint(i)
		// int8    -128..0..127
		// uint8   0..255
	}
	tmp := append([]uint(nil), arr...)
	fmt.Println(tmp)
}
