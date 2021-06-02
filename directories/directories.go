/*

by: @jurajstefanic | @stefanicjuraj

- useful functions for working with directories in the file system

*/

package main

// import
import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

// function check
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// function main
func main() {

	// create a new sub-directory in the current working directory
    err := os.Mkdir("subdir", 0755)
    check(err)

	// defer the removal - os.RemoveAll will delete a whole directory tree (similarly to rm -rf)
    defer os.RemoveAll("subdir")

	// function to create a new empty file
    createEmptyFile := func(name string) {
        d := []byte("")
        check(ioutil.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1")

	// create a hierarchy of directories - parents with MkdirAll
    err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)

	// create empty file
    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists directory contents
    c, err := ioutil.ReadDir("subdir/parent")
    check(err)

    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

	// Chdir lets us change the current working directory
    err = os.Chdir("subdir/parent/child")
    check(err)

	// see the contents of subdir/parent/child when listing the current directory
    c, err = ioutil.ReadDir(".")
    check(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

	// cd back
    err = os.Chdir("../../..")
    check(err)

	// visit a directory recursively
    fmt.Println("Visiting subdir")
    err = filepath.Walk("subdir", visit) // Walk accepts a callback function to handle every file or directory visited

} // end main

// function visit - called for every file or directory found recursively by filepath.Walk
func visit(p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", p, info.IsDir())
    return nil
} // end function visit