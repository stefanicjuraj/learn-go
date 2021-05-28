/*

by: @jurajstefanic | @stefanicjuraj

- you can have an if statement without an else
- a statement can precede conditionals; any variables declared in this statement are available in all branches
- no inline if (iff) - ternary if

*/

package main

// import
import (
	"fmt"
)

// function main
func main() {

	//if - else method
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

	// if method
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

	// if - else if - else method
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
    
} // end main