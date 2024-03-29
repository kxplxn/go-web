package _040502_gobPackage_

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Person struct {
	FirstName string
	LastName  string
}

func writeBinaryFile(data interface{}, file string) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(data)
	ioutil.WriteFile(file, buf.Bytes(), 0600)
}

func readBinaryFile(data interface{}, file string) {
	raw, _ := ioutil.ReadFile(file)
	buf := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buf)
	decoder.Decode(data)
}

func Demo() {
	fmt.Println("\n030502 Data Storage: Working With Gobs")

	file := "05-dataStorage/02-gobs/person"

	person := Person{"Bob", "Barker"}
	writeBinaryFile(person, file)

	var readPerson Person
	readBinaryFile(&readPerson, file)

	fmt.Printf("%+v\n", readPerson)
}
