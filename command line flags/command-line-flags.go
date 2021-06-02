/*

by: @jurajstefanic | @stefanicjuraj

- common way to specify options for command-line programs
- flag package supporting basic command-line flag parsing

*/

package main

// import
import (
    "flag"
    "fmt"
)

// function main
func main() {

	// string flag word with a default value "foo" and a short description - flag.String function returns a string pointer
    wordPtr := flag.String("word", "foo", "a string")

	// declares numb and fork flags - similar approach to the word flag
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

	// declare an option that uses an existing var declared elsewhere in the program - pass in a pointer to the flag declaration function
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

	// call flag.Parse() to execute the command-line parsing
    flag.Parse()

	// print - dump out the parsed options and any trailing positional arguments
    fmt.Println("word:", *wordPtr) // dereference the pointers with e.g. *wordPtr to get the actual option values
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())

} // end main