/*

by: @jurajstefanic | @stefanicjuraj

- lightweight thread of execution
- two function calls are running asynchronously in separate goroutines

*/

package main

// import
import (
	"fmt"
	"time"
)

// function f string
func f(from string) {

	// for method
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}

} // end function f string

// function main
func main() {
	
	f("direct")
	go f("goroutine") // invoke this function in a goroutine - use go f(s)

	// start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg) // print msg
	}("going")

	time.Sleep(time.Second) // sleep
	fmt.Println("finished") // print

} // end main