package main

import (
	"fmt"
)

func main() {
	//statePopulations := make(map[string]int, 10)
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)
    statePopulations["Georgia"] = 10454582
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Ohio"])
    delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Georgia"])
}
