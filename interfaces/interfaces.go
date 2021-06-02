/*

by: @jurajstefanic | @stefanicjuraj

- named collections of method signatures
- implement all the methods in the interface
- a variable has an interface type - call methods that are in the named interface
- circle and rect struct types both implement the geometry interface - use instances of these structs as arguments to measure

*/

package main

// import
import (
	"fmt"
	"math"
)

// basic interface for geometric shapes
type geometry interface {
	// attributes
	area() float64
	perim() float64
}

// implement this interface on rect
type rect struct {
	// attributes
	width, height float64
}

// implement this interface on circle types
type circle struct {
	// attribute radius
	radius float64
}

// implementation of rect with area
func (r rect) area() float64 {
	return r.width * r.height // return width x height
}

// implementation with area
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height // return 2*width x 2*width
}

// implementation for circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius // return pi x radius x radius
}

// implementation for circles
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius // return 2 x pi x radius
}

// function measure geometry
func measure(g geometry) {

	fmt.Println(g) // print g
	fmt.Println(g.area()) // print area g
	fmt.Println(g.perim()) // print perim g

} // end function measure

// function main
func main() {

	// assign
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// measure metod
	measure(r)
	measure(c)

} // end main