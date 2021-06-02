/*

by: @jurajstefanic | @stefanicjuraj

- ensure that a function call is performed later in a program’s execution - purposes of cleanup
- often used where e.g. ensure and finally would be used in other languages

*/

package main

// import
import (
	"fmt"
	"os"
)

// function main
func main() {

	// executed at the end of the enclosing function (main) - fter writeFile has finished
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)

} // end main

// function createFile
func createFile(p string) *os.File {

	fmt.Println("creating file") // print creating file
	f, err := os.Create(p)
	// if method
	if err != nil {
		panic(err)
	}
	return f
} // end fuction createFile

// function writeFile
func writeFile(f *os.File) {

	fmt.Println("closing file") // print closing file
	err := f.Close()
	// if method
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err) // print error
	}

} // end function writeFile

// function closeFile
func closeFile(f *os.File) {

	fmt.Println("closing file")
	err := f.Close()
	// if method
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err) // print error
		os.Exit(1)
	}

} // end fuction closeFile