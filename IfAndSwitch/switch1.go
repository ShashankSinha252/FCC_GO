package main

import (
	"fmt"
)

func main() {
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 3, 9:
		fmt.Println("two, three or nine")
	default:
		fmt.Println("different number")
	}
}
