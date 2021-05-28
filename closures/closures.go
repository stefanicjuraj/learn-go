/*

by: @jurajstefanic | @stefanicjuraj

- useful when you want to define a function inline without having to name it
- function intSeq returns another function, which we define anonymously in the body of intSeq
- call intSeq, assigning the result (a function) to nextInt
- function value captures its own i value - updated each time we call nextInt

*/

package main

// import
import (
	"fmt"
)

// function intSeq 
func intSeq() func() int {

	i := 0

	return func() int {

		i++
		return i

	}
} // end intSeq

// function main
func main() {

	nextInt := intSeq()

	fmt.Println(nextInt()) // print 
	fmt.Println(nextInt()) // print
	fmt.Println(nextInt()) // print

	newInts := intSeq() // new intSeq
	fmt.Println(newInts()) // print

} // end main