/*

by: @jurajstefanic | @stefanicjuraj

- variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls
- var declares 1 or more variables
- you can declare multiple variables at once
- variables declared without a corresponding initialization are zero-valued
- the := syntax is shorthand for declaring and initializing a variable

*/

package main

// import
import (
	"fmt"
)

// function main
func main() {

	// var a
	var a = "first var"
	fmt.Println(a)

	// var b and c
	var b, c = 1, 2
	fmt.Println(b, c)

	// var d
	var d = true
	fmt.Println(d)

	// var e int
	var e int
	fmt.Println(e)

	// print
	f := "test"
	fmt.Println(f)

} // end main
