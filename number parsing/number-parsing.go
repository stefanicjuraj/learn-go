/*

by: @jurajstefanic | @stefanicjuraj

- parsing numbers from strings is a basic but common task in many programs
- built-in package strconv provides the number parsing
- parse functions return an error on bad input

*/

package main

// import
import (
    "fmt"
    "strconv"
)

// function main
func main() {

    f, _ := strconv.ParseFloat("1.234", 64) // ParseFloat, this 64 tells how many bits of precision to parse
    fmt.Println(f) // print f

    i, _ := strconv.ParseInt("123", 0, 64) // ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits
    fmt.Println(i) // print i

    d, _ := strconv.ParseInt("0x1c8", 0, 64) // ParseInt will recognize hex-formatted numbers
    fmt.Println(d) // print d

    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u) // print u

    k, _ := strconv.Atoi("135") // Atoi is a convenience function for basic base-10 int parsing
    fmt.Println(k) // print k

    _, e := strconv.Atoi("error")
    fmt.Println(e) // print e

} // end main