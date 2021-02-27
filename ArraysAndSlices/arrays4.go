package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c := [...]int{7, 8, 9}
	d := &c
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)
}
