/*

by: @jurajstefanic | @stefanicjuraj

- slices are a key data type giving a more powerful interface to sequences than arrays
- composed into multi-dimensional data structures
- slices are typed only by the elements they contain (not the number of elements)
- declare and initialize a variable for slice in a single line
- to create an empty slice with non-zero length, use the built-in make
- len returns the length of the slice
-  slices support several more that make them richer than arrays:
	- append - returns a slice containing one or more new values
	- copy -  
	- slice - 

*/

package main

// import
import (
	"fmt"
)

// function main
func main() {

    s := make([]string, 3)
    fmt.Println("emp:", s) // print emp

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s) // print set
    fmt.Println("get:", s[2]) // print get

    fmt.Println("len:", len(s)) // print length

    s = append(s, "d") // append "d"
    s = append(s, "e", "f") // append "e" and "f"
    fmt.Println("apd:", s) // print append

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c) // print cpy

    l := s[2:5]
    fmt.Println("sl1:", l) // print slice 1

    l = s[:5]
    fmt.Println("sl2:", l) // print slice 2

    l = s[2:]
    fmt.Println("sl3:", l) // print slice 3

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t) // print dcl

    twoD := make([][]int, 3)
	// for method
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
		// for method
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }

    fmt.Println("2d: ", twoD) // print 2d
    
} // end main