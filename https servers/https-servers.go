/*

by: @jurajstefanic | @stefanicjuraj

- basic HTTP server using the net/http package
- handler is an object implementing the http.Handler interface
- using the http.HandlerFunc adapter on functions with the appropriate signature
- functions serving as handlers take a http.ResponseWriter and a http.Request as arguments
- response writer is used to fill in the HTTP response
- handler reads all the HTTP request headers and echos them into the response body
- registers handlers on server routes using the http.HandleFunc convenience function - sets up the default router in the net/http package and takes a function as an argument
- run the server in the background - access the /hello route

*/

package main

// import
import (
    "fmt"
    "net/http"
)

// function hello
func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n") // print
}

// function headers
func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h) // print
        } // end range
    } // end for
} // end function headers

// function main
func main() {

	// handle functions
    http.HandleFunc("/hello", hello) // hello
    http.HandleFunc("/headers", headers) // headers

	// call the ListenAndServe with the port and a handler
    http.ListenAndServe(":8090", nil) // nil tells it to use the default router we’ve just set up

} // end main