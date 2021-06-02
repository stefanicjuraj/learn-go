/*

by: @jurajstefanic | @stefanicjuraj

- represent a single event in the future
-  tell the timer how long you want to wait - provides a channel that will be notified at that time


*/

package main

// import
import (
	"fmt"
	"time"
)

// function main
func main() {

	// attribute timer new
	timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C blocks timer’s channel C until it sends a value indicating that the timer started
	<-timer1.C 
	fmt.Println("timer 1 start") // print timer

	// timer 2
	timer2 := time.NewTimer(time.Second)

	// function go - you can cancel the timer before it fires
	go func() {
		<-timer2.C
		fmt.Println("timer 2 start") // print timer 2
	}()
	
	// attribute stop
	stop2 := timer2.Stop()
	
	// if method
	if stop2 {
		fmt.Println("timer 2 stop") // print stop timer
	}

	time.Sleep(2 * time.Second)

} // end main