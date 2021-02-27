package main

import (
	"fmt"
)

func main() {
	n := 3.14
	n = 13.7e202
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)
}
