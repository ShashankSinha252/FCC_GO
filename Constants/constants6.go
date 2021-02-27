package main

import (
	"fmt"
)

const (
    _ = iota + 5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	//var specialistType int
	var specialistType int = dogSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)
    fmt.Printf("Cat specialist: %v\n", catSpecialist)
}
