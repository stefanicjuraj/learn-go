/*

by: @jurajstefanic | @stefanicjuraj

*/

package main

// import
import (
	"fmt"
	"time"
)

// function worker - run several concurrent instances
func worker(id int, jobs <-chan int, results chan<- int) {

	// for method - receive work on the jobs channel and send the corresponding results on results
	for j := range jobs {

		fmt.Println("worker ", id, "started job ", j) // print worker start job
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished job ", j) // print worker finish job
		results <- j * 2

	} // end for method

} // end function worker

// function main
func main() {

	// const
	const numJobs = 5
	// attributes - send work and collect their results
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// for method
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// for method - 5 jobs and close channel to indicate all the work
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // close method

	// for method - collect all the results of the work - worker goroutines have finished
	for a := 1; a <= numJobs; a++ {
		<-results
	}

} // end main