package main

import (
	"fmt"
	"time"
)

//Sample function to check race flag
func main() {
	var msg = "Ola"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Tata"
	time.Sleep(100 * time.Millisecond)
}
