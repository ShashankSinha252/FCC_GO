package main

import (
	"fmt"
)

func main() {
	greeting := "Fluffy"
	name := "boi"
	sayGreeting(&greeting, &name)
	fmt.Println(name)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
    *name = "cat"
    fmt.Println(*name)
}
