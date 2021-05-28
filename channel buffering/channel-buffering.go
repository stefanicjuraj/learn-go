/*

by: @jurajstefanic | @stefanicjuraj

- channels are unbuffered - only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value
- accept a limited number of values without a corresponding receiver for those values

*/

package main

// import
import (
	"fmt"
)

// function main
func main() {

	// channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	// send these values into the channel without a corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages) // print first message
	fmt.Println(<-messages) // print second message

} // end main