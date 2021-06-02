/*

by: @jurajstefanic | @stefanicjuraj

- for more complex state we can use a mutex to safely access data across multiple goroutines 
- running the program shows that we executed about 90,000 total operations against our mutex-synchronized state

*/

package main

// import
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// function main
func main() {

	// attributes
	var state = make(map[int]int) // state is a map
	var mutex = &sync.Mutex{} // syncronizes access to state
	var readOps uint64 // read option
	var writeOps uint64 // write option

	// for method - 100 goroutines - repeated reads
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock() // locking the mutex allows exclusive access to state
				total += state[key]
				mutex.Unlock() // unlock
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond) // wait for second read
			}
		}()
	}

	// for method - 10 goroutines - simulating writes with same pattern
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock() // lock
                state[key] = val
                mutex.Unlock() // unlock
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond) // wait
			}
		}()
	}

				time.Sleep(time.Second) // wait - let the 10 goroutines work on the state and mutex for a second
				
				// gets and reports final operation count
				readOpsFinal := atomic.LoadUint64(&readOps)
				fmt.Println("readOps: ", readOpsFinal) // print readops
				writeOpsFinal := atomic.LoadUint64(&writeOps)
				fmt.Println("writeOps: ", writeOpsFinal) // print writeops

				mutex.Lock() // lock
				fmt.Println("state: ", state) // print state
				mutex.Unlock() // unlock
				
} // end main