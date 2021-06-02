/*

by: @jurajstefanic | @stefanicjuraj

- filepath package provides functions to parse and construct file paths in a way that is portable between operating systems
- dir/file on Linux vs. dir\file on Windows
- use Join instead of concatenating /s or \s manually
- Join will also normalize paths by removing superfluous separators and directory changes

*/

package main

// import
import (
    "fmt"
    "path/filepath"
    "strings"
)

// function main
func main() {

    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)

	// Join should be used to construct paths in a portable way - takes any number of arguments and constructs a hierarchical path from them
    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir and Base can be used to split a path to the directory and the file - Split will return both in the same call
    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

	// check whether a path is absolute
    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

    filename := "config.json"

	// file names have extensions following a dot - split the extension out of such names with Ext
    ext := filepath.Ext(filename)
    fmt.Println(ext)

	// find the file’s name with the extension removed, use strings.TrimSuffix
    fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a base and a target - returns an error if the target cannot be made relative to base
    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel) // print rel

    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel) // print rel

} // end main