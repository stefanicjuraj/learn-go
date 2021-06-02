/*

by: @jurajstefanic | @stefanicjuraj

- common type of program that reads input on stdin, processes it, and then prints some derived result to stdout
- grep and sed are common line filters

*/

package main

// import
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// function main
func main() {

	// unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method that advances the scanner to the next token
	scanner := bufio.NewScanner(os.Stdin)

	// for method
	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text()) // Text returns the current token

		fmt.Println(ucl) // Write out the uppercased line

	} // end for

	// check for errors during Scan - end of file is expected and not reported by scan as an error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err) // print
		os.Exit(1) // exit
	}

} // end main
