package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	ms.foo = 42
	fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}
