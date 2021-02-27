package main

import (
	"fmt"
)

const a int = 14

func main() {
	const a = 27
	const b int16 = 32
	const c int = 29
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", a+b, a+b)
	fmt.Printf("%v, %T\n", a+c, a+c)
}
