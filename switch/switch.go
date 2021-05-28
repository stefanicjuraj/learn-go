/*

by: @jurajstefanic | @stefanicjuraj

- switch statements express conditionals across many branches
- switch without an expression is an alternate way to express if/else logic
- you can use commas to separate multiple expressions in the same case statement
- use the optional default case
- type switch compares types instead of values

*/

package main

// import
import (
	"fmt"
	"time"
)

// function main
func main() {

	i := 2
	fmt.Print("Write ", i , " as ") // print i

	// switch method
	switch i {
	case 1:
		fmt.Println("one") // print one
	case 2:
		fmt.Println("two") // print two
	case 3: 
		fmt.Println("three") // print three
	}

	// switch method
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is weekend!") // print weekend
	default:
		fmt.Println("it is weekday") // print weekday
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it is before noon") // print before noon
	default:
		fmt.Println("it is after noon") // print after noon
	}

	whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool") // print bool
        case int:
            fmt.Println("I'm an int") // print int
        default:
            fmt.Printf("Don't know type %T\n", t) // print dont know
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}