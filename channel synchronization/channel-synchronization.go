/*

by: @jurajstefanic | @stefanicjuraj

- use channels to synchronize execution across goroutines
- remove the <- done line from this program - the program would exit before the worker even started

*/

package main

// import
import (
	"fmt"
	"time"
)

// function worker - done channel will be used to notify another goroutine that this function’s work is done
func worker(done chan bool) {

	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Print("finished")

	done <- true // send a value to notify that done
}

// function main
func main() {

	// start worker goroutine - channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	<- done // block until received a notification from worker on the channel

} // end main