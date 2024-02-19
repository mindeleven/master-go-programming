/////////////////////////////////
// Coding along with the Udemy course
// Go (Golang) Programming: The Complete Go Bootcamp 2024
// https://www.udemy.com/course/master-go-programming-complete-golang-bootcamp/
// Source code and comments take from the videos of this course
/////////////////////////////////

package main

import "fmt"

func main() {
	SayHello()
}
func SayHello() {
	fmt.Println("Hello package v1.0.0!")

	var x int = 10
	_ = x
}
