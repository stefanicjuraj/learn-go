/*

by: @jurajstefanic | @stefanicjuraj

- allowed to convert a slice in the uppercase - ToUpper() function
- returns a copy of the given slice of bytes
- defined under the bytes package -  import bytes package for access ToUpper function

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
	slice_1 := []byte{'j', 'u', 'r', 'a', 'j'}
    slice_2 := []byte{'s', 't', 'e', 'f', 'a', 'n', 'i', 'c'}

	// display slice
	fmt.Printf("Slice 1: %s", slice_1)
    fmt.Printf("\nSlice 2: %s", slice_2)

	// convert to uppercase - ToUpper
	res1 := bytes.ToUpper(slice_1)
    res2 := bytes.ToUpper(slice_2)

	// display resulsts
	fmt.Printf("\nSlice 1: %s", res1)
    fmt.Printf("\nSlice 2: %s", res2)

} // end main