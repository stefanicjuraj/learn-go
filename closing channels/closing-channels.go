/*

by: @jurajstefanic | @stefanicjuraj

- indicates that no more values will be sent on it
- useful to communicate completion to the channel’s receivers

*/

package main

// import
import (
	"fmt"
)

// function main - use a jobs channel to communicate work to be done from the main() goroutine to a worker goroutine
func main() {

	// attributes
	jobs := make(chan int, 5)
	done := make(chan bool)

	// function go - repeatedly receives from jobs with j, more := <-jobs
	go func() {

		// for method
		for {
			j, more := <- jobs
			if more {
				fmt.Println("job received: ", j) // print job received
			} else {
				fmt.Println("all jobs received!") // print jobs received
				done <- true
				return
			} // end else
		} // end for

	}() // end function go

	// for method - sends 3 jobs to the worker over the jobs channel, then closes it
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("job sent", j) // print job sent
	} // end for

	close(jobs)
	fmt.Println("all jobs sent!") // print jobs sent

	<-done // await the worker using the synchronization approach

} // end main