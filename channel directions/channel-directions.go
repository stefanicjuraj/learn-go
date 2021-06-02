/*

by: @jurajstefanic | @stefanicjuraj

- specify if a channel is meant to only send or receive values
- specificity increases the type-safety of the program

*/

package main

// import
import (
	"fmt"
)

// function ping - only accepts a channel for sending values
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// function pong - accepts one channel for receives (pings) and a second for sends (pongs)
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

// function main
func main() {

    pings := make(chan string, 1)
    pongs := make(chan string, 1)

	// assign
    ping(pings, "passed message")
    pong(pings, pongs)

    fmt.Println(<-pongs) // print 

} // end main