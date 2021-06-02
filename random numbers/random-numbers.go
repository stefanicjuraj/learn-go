/*

by: @jurajstefanic | @stefanicjuraj

- math/rand package provides pseudorandom number generation
- default number generator is deterministic - produce the same sequence of numbers each time by default
- to produce varying sequences, give it a seed that changes
- not safe to use for random numbers you intend to be secret - use crypto/rand for those

*/

package main

// import
import (
    "fmt"
    "math/rand"
    "time"
)

// function main
func main() {

	// rand.Intn returns a random int n, 0 <= n < 100
    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))

    fmt.Println()

    fmt.Println(rand.Float64()) // rand.Float64 returns a float64 f, 0.0 <= f < 1.0

	// generate random floats in other ranges, for example 5.0 <= f' < 10.0
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)

    fmt.Println()

    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))

    fmt.Println()

    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))

    fmt.Println()
	
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))

} // end main


