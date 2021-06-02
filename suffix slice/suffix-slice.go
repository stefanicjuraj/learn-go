/*

by: @jurajstefanic | @stefanicjuraj

- slice is a variable-length sequence which stores elements of a similar type
- not allowed to store different type of elements in the same slice
- check whether the given slice ends with the specified prefix - HasSuffix() function
- return true if the given slice starts with the specified suffix
- return false if the given slice does not end with the specified suffix

*/

package main

// import
import (
	"fmt"
	"bytes"
)

// function main
func main() {

	// creating slice of bytes
	slice_1 := []byte{'G', 'E', 'E', 'K', 'S', 'F','o', 'r', 'G', 'E', 'E', 'K', 'S'}
    slice_2 := []byte{'A', 'P', 'P', 'L', 'E'}

	// checking whether the given slice ends with specified suffix
	res1 := bytes.HasSuffix(slice_1, []byte("S"))
    res2 := bytes.HasSuffix(slice_2, []byte("O"))
    res3 := bytes.HasSuffix([]byte("Hello! I am Apple."), []byte("Hello"))
    res4 := bytes.HasSuffix([]byte("Hello! I am Apple."), []byte("Apple."))
    res5 := bytes.HasSuffix([]byte("Hello! I am Apple."), []byte("."))
    res6 := bytes.HasSuffix([]byte("Hello! I am Apple."), []byte("P"))
    res7 := bytes.HasSuffix([]byte("Geeks"), []byte(""))

	// print & display results
	fmt.Println("Result_1: ", res1)
    fmt.Println("Result_2: ", res2)
    fmt.Println("Result_3: ", res3)
    fmt.Println("Result_4: ", res4)
    fmt.Println("Result_5: ", res5)
    fmt.Println("Result_6: ", res6)
    fmt.Println("Result_7: ", res7)

} // end main