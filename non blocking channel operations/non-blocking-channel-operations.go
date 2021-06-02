/*

by: @jurajstefanic | @stefanicjuraj

- sends and receives on channels that are blocking
- we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects

*/

package main

// import
import (
	"fmt"
)

// function main
func main() {

	// attributes
	messages := make(chan string)
	signals := make(chan bool)

	// select method - non-blocking receive -  a value is available on messages then select will take the <-messages case with that value - if not then default
	select {
	case msg := <-messages:
		fmt.Println("message received: ", msg)
	default:
		fmt.Println("message received null")
	}

	msg := "hi" 

	// select method -  msg cannot be sent to the messages channel - no buffer and no receiver
	select {
	case messages <- msg:
		fmt.Println("sent message: ", msg)
	default:
		fmt.Println("message sent null")
	}

	// select method - use multiple cases above the default clause to implement a multi-way non-blocking select - non-blocking receives on both messages and signals
	select {
	case msg := <-messages:
		fmt.Println("message received: ", msg)
	case sig := <-signals:
		fmt.Println("signal received: ", sig)
	default: 
		fmt.Println("activity null")
	}

} // end main