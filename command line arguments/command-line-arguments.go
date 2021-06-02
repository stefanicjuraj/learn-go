/*

by: @jurajstefanic | @stefanicjuraj

- common way to parameterize execution of programs
- go run hello.go uses run and hello.go arguments to the go program

*/

package main

// import
import (
    "fmt"
    "os"
)

// function main
func main() {

	// os.Args provides access to raw command-line arguments
    argsWithProg := os.Args
	// os.Args[1:] holds the arguments to the program
    argsWithoutProg := os.Args[1:]
	// get individual args with normal indexing
    arg := os.Args[3]

	// print
    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)

} // end main