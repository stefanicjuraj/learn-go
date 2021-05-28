/*

by: @jurajstefanic | @stefanicjuraj

- fact function calls itself until it reaches the base case of fact(0)

*/

package main

// import
import (
	"fmt"
)

// function fact
func fact(n int) int {

	if n == 0 {
		return 1 // return 1
	}
	return n * fact(n-1) // return n * fact(n-1)
}

// function main
func main() {

	fmt.Println(fact(5)) //print fact * 5

}