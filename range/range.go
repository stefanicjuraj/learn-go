/*

by: @jurajstefanic | @stefanicjuraj

- range iterates over elements in a variety of data structures
- range on arrays and slices provides both the index and value for each entry
- range on map iterates over key/value pairs
- range can also iterate over just the keys of a map
- range on strings iterates over Unicode code points
- first value is the starting byte index of the rune and the second the rune itself

*/

package main

// import
import (
	"fmt"
)

// function main
func main() {

    nums := []int{2, 3, 4} // values
    sum := 0 // value

	// for method
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum) // print sum

	// for method
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i) // print i
        }
    }

	// map
    kvs := map[string]string{"a": "apple", "b": "banana"}

	// for method
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v) // print k and v
    }

	// for method
    for k := range kvs {
        fmt.Println("key:", k) // print key k
    }

	// for method
    for i, c := range "go" {
        fmt.Println(i, c) // print i and c
    }

} // end main