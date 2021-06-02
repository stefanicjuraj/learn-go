/*

by: @jurajstefanic | @stefanicjuraj

- allows you to pass references to values and records within your program
- zeroval has an int parameter - arguments will be passed by value
- zeroval will get a copy of ival distinct from the one in the calling function
- zeroval doesn’t change the i in main
- zeroptr in contrast has an *int parameter - takes an int pointer
- zeroptr changes the i in main - a reference to the memory address for that variable
- *iptr code in the function body dereferences the pointer from its memory address to the current value at that address
- assigning a value to a dereferenced pointer changes the value at the referenced address
- The &i syntax gives the memory address of i, i.e. a pointer to i
- pointers can be printed

*/

package main

// import
import (
	"fmt"
)

// function zeroval
func zeroval (ival int) {
	ival = 0 // value 0
}

// function zeroptr
func zeroptr (iptr *int) {
	*iptr = 0 // value 0
}

// function main
func main() {

	i := 1 // value of i

	fmt.Println("initial: ", i) // print initial + i

	zeroval(i)
	fmt.Println("zeroval: ", i) // print zeroval + i

	zeroptr(&i)
	fmt.Println("zeroptr: ", i) // print zeroptr + i

	fmt.Println("pointer: ", &i) // print pointer to &i

} // end main