package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	//b := Bird{}
	//b.Name = "Emu"
	//b.Origin = "Australia"
	//b.SpeedKPH = 48
	//b.CanFly = false
	fmt.Println(b)
}
