/////////////////////////////////
// The Typical Structure of a Go Application
// Coding along with the Udemy course
// Go (Golang) Programming: The Complete Go Bootcamp 2024
// https://www.udemy.com/course/master-go-programming-complete-golang-bootcamp/
// Source code and comments take from the videos of this course
/////////////////////////////////

// a package clause starts every source file
// main is a special name declaring an executable rather than a library (package)
package main

// import declaration declares packages used in this file
// https://pkg.go.dev/fmt
import "fmt"

// package scoped variables and constants
var x int = 100

const secondsInHour = 3600

// a function declaration
// main is a special function name
// it is the entry point for the executable program
func main() {
	// Local Scoped Variables and Constants Declarations, calling functions etc
	var a int = 7
	var b float64 = 3.5
	const c int = 10

	fmt.Println("Hello Go World!!")
	fmt.Printf("An hour has %b seconds \n", secondsInHour)
	distance := 60.8 // distance in km
	fmt.Printf("The distance in miles is %f \n", distance*0.621)
	fmt.Println(a, b, c, x) // => 7 3.5 10 100
}
