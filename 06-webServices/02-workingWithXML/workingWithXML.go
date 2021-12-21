package _030602_workingWithXML_

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	FirstName string   `xml:"fname""`
	LastName  string   `xml:"lname""`
	Hobbies   []string `xml:"hobbies""`
}

func Demo() {
	fmt.Println("\n030602 Web Services: Working with XML")

	joe := Person{FirstName: "Joe", LastName: "Smith"}
	joe.Hobbies = []string{"Skiing", "Wind Surfing"}

	xmlOutput, _ := xml.MarshalIndent(&joe, " ", "   ")
	fmt.Print(xml.Header)
	fmt.Println(string(xmlOutput))

	joeFromXML := Person{}
	xml.Unmarshal(xmlOutput, &joeFromXML)
	fmt.Println(joeFromXML)
}
