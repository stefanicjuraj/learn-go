/*

by: @jurajstefanic | @stefanicjuraj

- built-in support for JSON encoding and decoding
- to and from built-in and custom data types

*/

package main

// import
import (
	"encoding/json"
	"fmt"
	"os"
)

//  two structs to demonstrate encoding and decoding of custom types

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

// function main
func main() {

	// encoding basic data types to JSON strings
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB)) // print string

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB)) // print string

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB)) // print string

	// encode to JSON arrays and objects
    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB)) // print string

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB)) // print string

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB)) // print map

	// automatically encode your custom data types - include exported fields in the encoded output and will by default use those names as the JSON keys
    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B)) // print string

	// use tags on struct field declarations to customize the encoded JSON key names
    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B)) // print string

	// decoding JSON data into Go values
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    var dat map[string]interface{} // provide a variable where the JSON package can put the decoded data - hold a map of strings to arbitrary data types

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }

    fmt.Println(dat) // print dat

	// convert them to their appropriate type - convert the value in num to the expected float64 type
    num := dat["num"].(float64)
    fmt.Println(num) // print num

	// nested data requires a series of conversions
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1) // print str1

	// decode JSON into custom data types - advantages of adding type-safety to programs and eliminating type assertions when accessing the decoded data
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res) // print res
    fmt.Println(res.Fruits[0]) // print res fruits

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)

} // end main