package main

import (
	"fmt"
)

func main() {
	sp := map[string]int{
		"Texas":    34540991,
		"Florida":  11614374,
		"New York": 12802503,
	}
	if pop, ok := sp["Florida"]; ok {
		fmt.Println(pop)
	}
}
