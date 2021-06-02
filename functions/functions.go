/*

by: @jurajstefanic | @stefanicjuraj

- functions are central in Go
- Go requires explicit returns, i.e. it won’t automatically return the value of the last expression
- when you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type
- call a function - name(args)

*/

package main

import (
	"fmt"
)

// function plus
func plus(a int, b int) int {
    return a + b // return
}

// function plus plus
func plusPlus(a, b, c int) int {
    return a + b + c // return
}

// function main
func main() {

    res := plus(1, 2)
    fmt.Println("1+2 =", res) // print result

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res) // print result

} // end main