package main

import (
	"fmt"
	"math"
)

func main() {
	const myConst float = math.Sin(1.57)
	fmt.Printf("%v, %T\n", myConst, myConst)
}
