package main

import "fmt"

func main() {
	s := "Hallöchen"
	//   "Hall..chen"
	//    0123456789

	fmt.Println(s)
	fmt.Println(s[4:6])
	//        ^^
	fmt.Println(s[4:5])
	//        ^
	fmt.Println(s[:4])
	//    ^^^^
	fmt.Println(s[5:])
	//         ^^^^^
}
