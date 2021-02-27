package main

import (
	"fmt"
)

func main() {
	sp := map[string]int{
		"Texas":   43660981,
		"Florida": 98103327,
		"Ohio":    74012856,
	}
	fmt.Println(sp)

	pop, ok := sp["Oho"]
	fmt.Println(pop, ok)

	pop, ok = sp["Ohio"]
	fmt.Println(pop, ok)

    fmt.Printf("Length: %v\n", len(sp))
}
