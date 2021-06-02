/*

by: @jurajstefanic | @stefanicjuraj

- panic typically means something went unexpectedly wrong
- use it to fail fast on errors that shouldn’t occur during normal operation
- check for unexpected errors
- idiomatic to use error-indicating return values wherever possible

*/

package main

// import
import (
	"os"
)

// function main
func main() {

	panic("problem!") // print panic

	//an unexpected error when creating a new file
	_, err := os.Create("/tmp/file")

	if err != nil {
		panic(err)
	}

} // end main