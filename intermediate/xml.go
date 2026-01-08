package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Preson struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name,omitempty"`
	Age     int      `xml:"age,omitempty"`
	City    string   `xml:"city,omitempty"`
	Email   string   `xml:"email,omitempty"`
	Address Address  `xml:"address,omitempty"`
}

type Address struct {
	City  string `xml:"city,omitempty"`
	State string `xml:"state,omitempty"`
}

func main() {

	person := Preson{Name: "John", Email: "exm@x.x"}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println("XML Data:", string(xmlData))

	xmlData, err = xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println(string(xmlData))

	// xmlRaw := `<person><name>Max</name><age>44</age></person>`
	xmlRaw := `<person><name>Max</name><age>44</age><address><city>Las Vegas</city><state>NV</state></address></person>`
	var personXML Preson
	err = xml.Unmarshal([]byte(xmlRaw), &personXML)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println(personXML)
	fmt.Println("-----")
	fmt.Println(personXML.XMLName.Local)
	fmt.Println("-----")
	fmt.Println(personXML.XMLName.Space)
	fmt.Println("-----")

	book := Book{
		ISBN:    "432-512-344",
		Title:   "Go",
		Author:  "Ahhhh",
		Another: "xxxxxxxxxxxas",
	}

	xmlBookAttr, err := xml.MarshalIndent(book, "a", "  ")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println(string(xmlBookAttr))
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	ISBN    string   `xml:"isbn,attr"`
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author,attr"`
	Another string   `xml:"another"`
}
