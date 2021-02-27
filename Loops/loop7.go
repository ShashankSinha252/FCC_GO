package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	fmt.Println()
	t := "Hello bro!"
	for k, v := range t {
		fmt.Println(k, string(v))
	}

	fmt.Println()
	sp := map[string]int{
		"Texas":    23423425,
		"New York": 80345201,
	}
	for k, v := range sp {
		fmt.Println(k, v)
	}
}
