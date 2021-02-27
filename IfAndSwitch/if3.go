package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := -30
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("Guess must be between 1 and 100!")
	}
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
	fmt.Println(!false)
}

func returnTrue() bool {
	fmt.Println("Returning true")
	return true
}
