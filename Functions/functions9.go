package main

import (
	"fmt"
)

func main() {
	var f func() = func() {
		fmt.Println("Ahoy")
	}
	f()
}
