package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"Texas":    43865401,
		"Ohio":     23543951,
		"New York": 34563242,
	}
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}
