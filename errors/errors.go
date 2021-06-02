/*

by: @jurajstefanic | @stefanicjuraj

- idiomatic to communicate errors via an explicit, separate return value
- easy to see which functions return errors
- handle them using the same language constructs employed for any other, non-error tasks
- errors are the last return value and have type error - built-in interface

*/

package main

// import
import (
	"errors"
	"fmt"
)

// function f1
func f1(arg int) (int, error) {

	// if method
	if arg == 42 {
		return -1, errors.New("42 is not a valid number here") // constructs a basic error value with the given error message
	}
	return arg + 3, nil // return - nil value in the error position indicates no error
} // end function f1

// type struct
type argError struct {
	// attributes
	arg int
	prob string
}

// use custom types as errors by implementing the Error() method on them
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob) // return
}

// function f2
func f2(arg int) (int, error) {
	// if method
	if arg == 42 {
		return -1, &argError{arg, "not valid"} //  use &argError syntax to build a new struct - values for two fields arg and prob
	}
	return arg + 3, nil // return
}

// function main
func main() {

	// test out each of our error-returning functions
	for _, i := range []int{7, 42} {
		if r,e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e) // print fail
		} else {
			fmt.Println("f1 worked: ", r) // print success
		}
	}
	// test out each of our error-returning functions
	for _, i := range []int{7, 42} {
		if r,e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e) // print fail
		} else {
			fmt.Println("f2 worked: ", r) // print success
		}
	}

	// get the error as an instance of the custom error type via type assertion
	_, e := f2(42) // set / assign
	
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg) // print arg
		fmt.Println(ae.prob) // print prob
	}

} // end main