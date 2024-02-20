/////////////////////////////////
// Coding along with the Udemy course
// Go (Golang) Programming: The Complete Go Bootcamp 2024
// https://www.udemy.com/course/master-go-programming-complete-golang-bootcamp/
// Source code and comments take from the videos of this course
/////////////////////////////////

package main

import "fmt"

func main() {
	// declaring variables
	// (1) using the var keyword
	var age int = 30
	fmt.Println("My age is", age)

	var name = "Dan"
	fmt.Println("My name is", name)

	var name2 = "Alice"
	// underscore to mute "declared and not used" compiler error
	_ = name2

	// short declaration operator
	s := "Learning Golang!"
	fmt.Println(s)
}
