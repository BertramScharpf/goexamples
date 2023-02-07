package main

import "fmt"

type Car struct{}

type Limousine struct {
	Car
}

func (_ Car) consumption() float64 {
	return 0.07
}

func (_ Limousine) consumption() float64 {
	return 0.15
}

func main() {
	var c Car = Car{}
	var l Limousine = Limousine{}
	var p *Car = &l.Car

	fmt.Println("Verbrauch Auto:      ", c.consumption())
	fmt.Println("Verbrauch Limousine: ", l.consumption())
	fmt.Println("Falsch:              ", l.Car.consumption())
	fmt.Println("Zeiger:              ", p.consumption())
}
