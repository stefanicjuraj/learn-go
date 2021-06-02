/*

by: @jurajstefanic | @stefanicjuraj

- wait for multiple goroutines to finish
- order of workers starting up and finishing is likely to be different for each invocation

*/

package main

// import
import (
	"fmt"
	"sync"
	"time"
)

// function worker - WaitGroup must be passed to functions by pointer
func worker(id int, wg *sync.WaitGroup) {

	// notify the WaitGroup that we’re done
	defer wg.Done()

	fmt.Printf("worker %d starting\n", id) // print start

	time.Sleep(time.Second) // sleep method to simulate an expensive task

	fmt.Printf("worker %d done \n", id) // print done

} // end fuction worker

// function main
func main() {

	// used to wait for all the goroutines launched here to finish
	var wg sync.WaitGroup

	// for method - launch several goroutines and increment the WaitGroup counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait() // block until the WaitGroup counter goes back to 0; all the workers notified they’re done

} // end main