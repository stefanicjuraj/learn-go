/*

by: @jurajstefanic | @stefanicjuraj

- called with any number of trailing arguments
- called in the usual way with individual arguments
- fmt.Println is a common variadic function

*/

package main

// import
import (
	"fmt"
)

func sum(nums ...int) {

	fmt.Print(nums, " ")
	total := 0

	// for method
	for _, num := range nums {
		total += num
	}

	fmt.Println(total) // print total

} // end sum

// function main
func main() {

	sum(1, 2) // sum
	sum(1, 2, 3) // sum

	nums := []int{1, 2, 3, 4}
    sum(nums...)
	
} // end main