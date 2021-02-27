package main

import (
	"fmt"
)

func main() {
	var students [6]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[2] = "Ahmed"
	students[1] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Number of student : %v\n", len(students))
}
