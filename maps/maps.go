/*

by: @jurajstefanic | @stefanicjuraj

- built-in associative data type
- use the built-in make: make(map[key-type]val-type)
- set key/value pairs using typical name[key] = val syntax
- printing a map with e.g. fmt.Println will show all of its key/value pairs
- get a value for a key with name[key]
- the built-in len returns the number of key/value pairs when called on a map
- the built-in delete removes key/value pairs from a map
- note that maps appear in the form map[k:v k:v] when printed with fmt.Println

*/

package main

// import
import (
	"fmt"
)

// function main
func main() {

	// make map m
    m := make(map[string]int)

    m["k1"] = 7 // values
    m["k2"] = 13 // values

    fmt.Println("map:", m) // print map m

    v1 := m["k1"]
    fmt.Println("v1: ", v1) // print v1

    fmt.Println("len:", len(m)) // print length

    delete(m, "k2") // delete 
    fmt.Println("map:", m) // print map m

    _, prs := m["k2"]
    fmt.Println("prs:", prs) // print prs

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n) // print map n

} //end main