/*

by: @jurajstefanic | @stefanicjuraj

- universal mechanism for conveying configuration information to Unix programs
- list of keys in the environment will depend on your particular machine

*/

package main

// import
import (
    "fmt"
    "os"
    "strings"
)

// function main
func main() {

	// set a key/value pair - use os.Setenv
    os.Setenv("FOO", "1")
	//  get a value for a key - use os.Getenv - return an empty string if the key isn’t present in the environment
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

	// print
    fmt.Println()

	// os.Environ - list all key/value pairs in the environment - returns a slice of strings in the form KEY=value
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2) // strings.SplitN them to get the key and value
        fmt.Println(pair[0]) // print all the keys
    }

} // end main

