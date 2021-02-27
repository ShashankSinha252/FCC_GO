package main

import (
	"fmt"
)

func main() {
	s := "looks like a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}
