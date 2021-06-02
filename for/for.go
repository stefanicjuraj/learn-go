/*
by: @jurajstefanic | @stefanicjuraj

- for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function
- you can also continue to the next iteration of the loop

*/

package main

// import
import (
	"fmt"
)

// function main
func main() {

	i := 1 // set value for i

	// for method
	for i <= 3 {
		fmt.Println(i) //print i
		i = i + i
	}

	// for method
	for j := 7; j <= 9; j++ {
		fmt.Println(j) // print j
	}

	// for method
	for {
		fmt.Println("loop") // print loop
		break
	}

	// for method
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)  // print n
	}

} // end main