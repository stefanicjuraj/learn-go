/*

by: @jurajstefanic | @stefanicjuraj

- sequence of variable-width characters where each and every character is represented by one or more bytes using UTF-8 Encoding
- check whether the string begins with the specified prefix or not with the help of HasPrefix() function
- import strings package in your program for accessing HasPrefix function
- return true if the given string starts with the specified prefix
- return false if the given string does not start with the specified prefix

*/

package main

// import
import (
	"fmt"
	"strings"
)

// function main
func main() {

	// strings
	s1 := "this is a test sentence"
	s2 := "my name is Juraj"
	s3 := "i study IT"

	// check for prefix - function hasPrefix()
	res1 := strings.HasPrefix(s1, "this")
	res2 := strings.HasPrefix(s2, "my")
	res3 := strings.HasPrefix(s3, "i")

	// print & display results
	fmt.Println("result 1: ", res1)
	fmt.Println("result 2: ", res2)
	fmt.Println("result 3: ", res3)

} // end main