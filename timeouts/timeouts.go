/*

by: @jurajstefanic | @stefanicjuraj

- important for programs that connect to external resources or need to bound execution time
- easy and elegant thanks to channels and select

*/

package main

// import
import (
	"fmt"
	"time"
)

// function main
func main() {

	c1 := make(chan string, 1)
	// function
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// select implementing a timeout - select proceeds with the first receive that’s ready
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // after awaits a value to be sent after the timeout
		fmt.Println("timeout 1") // print timeout
	}

	c2 := make(chan string, 1)
	// function
	go func() {
		time.Sleep(2 * time.Second) // sleep
		c2 <- "result 2"
	}()

	// select implementing a timeout - select proceeds with the first receive that’s ready
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second): // after awaits a value to be sent after the timeout
		fmt.Println("timeout 2") // print timeout
	}

} // end main