package main

import (
	"fmt"
)

func main() {
	n := 1 + 2i
	n = complex(5, 12)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("Real: %v, %T\n", real(n), real(n))
	fmt.Printf("Imaginary: %v, %T\n", imag(n), imag(n))
}
