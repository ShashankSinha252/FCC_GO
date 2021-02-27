package main

import (
	"fmt"
)

func main() {
	number := 50
	//guess := -30
	//guess := 300
	guess := 300
	if guess < 1 {
		fmt.Println("Guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("Guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("Correct!")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
}
