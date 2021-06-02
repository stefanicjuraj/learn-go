/*

by: @jurajstefanic | @stefanicjuraj

- common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds

*/

package main

// import
import (
    "fmt"
    "time"
)

// function main
func main() {

	// time.Now with Unix or UnixNano to get elapsed time since the Unix epoch in seconds or nanoseconds
    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)

	// get the milliseconds - to manually divide from nanoseconds
    millis := nanos / 1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)

	// convert integer seconds or nanoseconds since the epoch into the corresponding time
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))

} // end main