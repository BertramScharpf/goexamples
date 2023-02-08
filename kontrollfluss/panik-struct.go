package main

import "fmt"

type SomeError struct {
	Msg string
}

func catchit() {
	if r := recover(); r != nil {
		fmt.Printf("Fehler: %v (%T)\n", r)
	} else {
		fmt.Println("Alles sauber.")
	}
}

func main() {
	defer catchit()
	panic(&SomeError{"Arrrrgh"})  // Bereits hier wird ein String daraus gemacht.
	fmt.Println("Nach Arrrrgh.")
}
