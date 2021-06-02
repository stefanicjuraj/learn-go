/*

by: @jurajstefanic | @stefanicjuraj

- supports constants of character, string, boolean, and numeric values
- const declares a constant value
- const statement can appear anywhere a var statement can

*/

package main

// import
import (
	"fmt"
	"math"
)

// const
const x string = "constant"

// function main
func main() {

	// print
	fmt.Println(x)

	// const
	const n = 50000000
	const d = 3e20 / n

	// print	
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}