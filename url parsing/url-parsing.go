/*

by: @jurajstefanic | @stefanicjuraj

- URLs provide a uniform way to locate resources

*/

package main

// import
import (
    "fmt"
    "net"
    "net/url"
)

// function main
func main() {

    s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// parse the URL - ensure there are no errors
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

	// accessing the scheme
    fmt.Println(u.Scheme)

	// user contains all authentication info; call Username and Password on this for individual values
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

	// Host contains both the hostname and the port, if present. Use SplitHostPort to extract them
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

	// extract the path and the fragment after the #
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

	// get query params in a string of k=v format - use RawQuery
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
	// print
    fmt.Println(m)
    fmt.Println(m["k"][0])

} // end main