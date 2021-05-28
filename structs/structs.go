/*

by: @jurajstefanic | @stefanicjuraj

- structs are typed as collections of fields - useful for grouping data together to form records
- structs are mutable

*/

package main

// import
import (
	"fmt"
)

// type person struct type has name and age fields
type person struct {
	// attributes
	name string
	age int
}

// function newPerson - 
func newPerson (name string) *person {

	// attributes
	p := person {name: name}
	p.age = 42
	return &p

} // end function newPerson

// function main
func main() {

	fmt.Println(person{"Juraj", 21}) // syntax creates a new struct

	fmt.Println(person{name: "John", age: 30}) // name the fields when initializing a struct

	fmt.Println(person{name: "Jack"}) // omitted fields will be zero-valued

	fmt.Println(&person{name: "Ana", age: 40}) // & prefix yields a pointer to the struct

	fmt.Println(newPerson("Alice")) // idiomatic to encapsulate new struct creation in constructor functions

	s := person{name: "Laura", age: 50} // access struct fields with a dot
	fmt.Println(s.name) // print name

	sp := &s // use dots with struct pointers - the pointers are automatically dereferenced
	fmt.Println(sp.age) // print age

	sp.age = 51 // assign new age
	fmt.Println(sp.age) // print new age

} // end main