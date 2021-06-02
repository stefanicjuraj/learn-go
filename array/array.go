/*

by: @jurajstefanic | @stefanicjuraj

- an array is a numbered sequence of elements of a specific length
- the type of elements and length are both part of the array’s type
- the built-in len returns the length of an array
- array types are one-dimensional, but you can compose types to build multi-dimensional data structures

*/

package main

// import
import (
	"fmt"
)

// function main
func main() {

	// var
	var a [5]int
	fmt.Println("emp:", a) //print emp

	a[4] = 100
	fmt.Println("set:", a)    // print set
	fmt.Println("get:", a[4]) // print get

	fmt.Println("len:", len(a)) // print length

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b) // print dcl

	// var twoD
	var twoD [2][3]int
	// for method
	for i := 0; i < 2; i++ {
		// for method
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // print 2d
}// end main
