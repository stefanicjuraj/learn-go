/*

by: @jurajstefanic | @stefanicjuraj

- Go has built-in support for multiple return values
- used often in idiomatic Go - return both result and error values from a function
- if you only want a subset of the returned values, use the blank identifier _

*/

package main

// import
import (
	"fmt"
)

// function vals
func vals() (int, int) {
    return 3, 7 // return
}

// function main
func main() {

    a, b := vals() // values
    fmt.Println(a) // print a
    fmt.Println(b) // print b

    _, c := vals() // value
    fmt.Println(c) // print set value (c)
}