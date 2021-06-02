/*

by: @jurajstefanic | @stefanicjuraj

- supports methods defined on struct types
- automatically handles conversion between values and pointers for method calls
- use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct

*/

package main

// import
import (
	"fmt"
)

// type 
type rect struct {
	width, height int
}

// function with area and int - area method has a receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height // return
}

// function with perim and int - defined for either pointer or value receiver types (in this case, value receiver)
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// function main
func main() {

	r := rect{width: 10, height: 5} // r attributes

	// call the 2 methods defined for our struct
	fmt.Println("area: ", r.area()) // print area
	fmt.Println("perim: ", r.perim()) // print perim

	rp := &r // assign

	fmt.Println("area :", rp.area()) // print area
	fmt.Println("perim: ", rp.perim()) // print perim

} // end main