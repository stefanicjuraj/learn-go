/*

by: @jurajstefanic | @stefanicjuraj

- selecting all items that satisfy a given predicate or mapping all items to a new collection with a custom function
- provide collection functions if and when they are specifically needed for your program and data types
- example collection functions for slices of strings
- anonymous functions - you can also use named functions of the correct type

*/

package main

// import
import (
    "fmt"
    "strings"
)

// function index - returns the first index of the target string t, or -1 if no match is found
func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

// function include - returns true if the target string t is in the slice
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
} // end function include

// function any - returns true if one of the strings in the slice satisfies the predicate f
func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
} // end function any

// function all - returns true if all of the strings in the slice satisfy the predicate f
func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
} // end function all

// function filter - returns a new slice containing all strings in the slice that satisfy the predicate f
func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
} // end function filter

// function map - returns a new slice containing the results of applying the function f to each string in the original slice
func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
} // end function map

// function main
func main() {

	// attribute - try out our various collection functions
    var strs = []string{"peach", "apple", "pear", "plum"}

    fmt.Println(Index(strs, "pear")) // print pear

    fmt.Println(Include(strs, "grape")) // print grape

    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p") // print with prefix
    }))

    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p") // print with prefix
    }))

    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e") // print with contain
    }))

    fmt.Println(Map(strs, strings.ToUpper)) // print uppercase

} // end main