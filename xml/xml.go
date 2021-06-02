/*

by: @jurajstefanic | @stefanicjuraj

- built-in support for XML and XML-like formats
- encoding.xml package
- if the XML is malformed or cannot be mapped - descriptive error will be returned

*/

package main

// import
import (
    "encoding/xml"
    "fmt"
)

// type struct
type Plant struct {
	// attributes
    XMLName xml.Name `xml:"plant"`
    Id      int      `xml:"id,attr"`
    Name    string   `xml:"name"`
    Origin  []string `xml:"origin"`
} // end type

// function string - mapped to XML - field tags contain directives for the encoder and decoder
func (p Plant) String() string {
    return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

// function main
func main() {

    coffee := &Plant{Id: 27, Name: "Coffee"}
    coffee.Origin = []string{"Ethiopia", "Brazil"}

	// emit XML representing our plant - MarshalIndent produces a more human-readable output
    out, _ := xml.MarshalIndent(coffee, " ", "  ")
    fmt.Println(string(out)) // print string

    fmt.Println(xml.Header + string(out)) // generic XML header to the output

    var p Plant

	// if method -  Unmarhshal to parse a stream of bytes with XML into a data structure
    if err := xml.Unmarshal(out, &p); err != nil {
        panic(err)
    }

    fmt.Println(p) // print p

    tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}

	// type nesting - parent>child>plant field tag tells the encoder to nest all plants under <parent><child>
    type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants  []*Plant `xml:"parent>child>plant"`
    }

    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}

    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out)) // print string

} // end main