/*

by: @jurajstefanic | @stefanicjuraj

- replace all the elements in the given slice - ReplaceAll() function
- used to replace all the elements of the old slice with a new slice

*/

package main

// import
import (
	"fmt"
	"bytes"
)

// function main
func main() {

	// create slice of bytes
	slice_1 := []byte{'J', 'U', 'R', 'A','J'}
	slice_2 := []byte{'S','T','E','F','A','N','I','C'}

	// display slice
	fmt.Printf("Slice 1: %s", slice_1)
    fmt.Printf("\nSlice 2: %s", slice_2)

	// replace bytes
	res1 := bytes.ReplaceAll(slice_1, []byte("J"), []byte("j"))
    res2 := bytes.ReplaceAll(slice_2, []byte("S"), []byte("s"))

	// display results
	fmt.Printf("\nSlice 1: %s", res1)
    fmt.Printf("\nSlice 2: %s", res2)

} // end main