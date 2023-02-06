package main

import "fmt"

func catchit() {
	if r := recover(); r != nil {
		fmt.Println("Fehler:", r)
	} else {
		fmt.Println("Alles sauber.")
	}
}

func main() {
	defer catchit()
	panic("Arrrrgh!")
}
