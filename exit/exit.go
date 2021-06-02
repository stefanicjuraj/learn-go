/*

by: @jurajstefanic | @stefanicjuraj

- use os.Exit to immediately exit with a given status
- does not use an integer return value from main to indicate exit status
- exit with a non-zero status - use os.Exit
- if you run exit.go using go run - the exit will be picked up by go and printed
- by building and executing a binary - see the status in the terminal
- NOTE: - ! from our program never got printed

*/

package main

// import
import (
    "fmt"
    "os"
)

// function main
func main() {

	// defers will not be run when using os.Exit - fmt.Println will never be called
    defer fmt.Println("!")

	// exit with status 3
    os.Exit(3)

} // end main