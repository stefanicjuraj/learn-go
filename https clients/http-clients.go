/*

by: @jurajstefanic | @stefanicjuraj

- support for HTTP clients and servers in the net/http package
- issue simple HTTP requests
- http.Get is a convenient shortcut around creating an http.Client object and calling its Get method - uses the http.DefaultClient object which has useful default settings

*/

package main

// import
import (
    "bufio"
    "fmt"
    "net/http"
)

// function main
func main() {

	// issue an HTTP GET request to a server
    resp, err := http.Get("http://gobyexample.com")
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

	// print the HTTP response status
    fmt.Println("Response status:", resp.Status)

	// print the first 5 lines of the response body
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

} // end main