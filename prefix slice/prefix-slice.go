/*

by: @jurajstefanic | @stefanicjuraj

- slice is a variable-length sequence which stores elements of a similar type
- not allowed to store different type of elements in the same slice
- check whether the given slice begins with the specified prefix - HasPrefix() function
- return true if the given slice starts with the specified prefix
- return false if the given slice does not start with the specified prefix

*/

package main

// import
import (
	"bytes"
	"fmt"
)

// function main
func main() {

	// creating slice of bytes
	slice_1 := []byte{'G', 'E', 'E', 'K', 'S', 'F','o', 'r', 'G', 'E', 'E', 'K', 'S'}
	slice_2 := []byte{'A', 'P', 'P', 'L', 'E'}

	// checking whether the given slice starts with specified prefix
	res1 := bytes.HasPrefix(slice_1, []byte("G"))
	res2 := bytes.HasPrefix(slice_2, []byte("O"))
	res3 := bytes.HasPrefix([]byte("Hello! I am Apple."), []byte("Hello"))
	res4 := bytes.HasPrefix([]byte("Hello! I am Apple."), []byte("Welcome"))
	res5 := bytes.HasPrefix([]byte("Hello! I am Apple."), []byte("H"))
	res6 := bytes.HasPrefix([]byte("Hello! I am Apple."), []byte("X"))
	res7 := bytes.HasPrefix([]byte("Geeks"), []byte(""))

	// print & display results
	fmt.Println("Result_1: ", res1)
	fmt.Println("Result_2: ", res2)
	fmt.Println("Result_3: ", res3)
	fmt.Println("Result_4: ", res4)
	fmt.Println("Result_5: ", res5)
	fmt.Println("Result_6: ", res6)
	fmt.Println("Result_7: ", res7)

} // end main
