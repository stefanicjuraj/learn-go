/*

by: @jurajstefanic | @stefanicjuraj

- channels are the pipes that connect concurrent goroutines
- send values into channels from one goroutine and receive those values into another goroutine
- typed by the values they convey
- <- channel syntax receives a value from the channel

*/

package main

// import
import (
	"fmt"
)

// function main
func main() {

	messages := make(chan string)

	// go function to send "ping" via messages
	go func() {
		messages <- "ping" // send a value into a channel using the channel <- syntax
	}()

	msg := <-messages
	fmt.Println(msg) // print msg

} // end main