/*

by: @jurajstefanic | @stefanicjuraj

- use this syntax to iterate over values received from a channel 

*/

package main

// import
import (
	"fmt"
)

// function main
func main() {

	// attributes - iterate over 2 values in the queue channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// close method
	close(queue)

	// for method - range iterates over each element as it’s received from queue - closed the channel above - iteration terminates after receiving the 2 elements
	for elem := range queue {
		fmt.Println(elem) // print element
	}

} // end main