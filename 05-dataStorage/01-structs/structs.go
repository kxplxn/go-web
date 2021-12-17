package _040501_structs_

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

type Group struct {
	Name   string
	People []Person
}

func Demo() {
	fmt.Println("\n030501 Data Storage: In-Memory Storage with Structs")

	person1 := Person{"George", "Smith"}
	person2 := Person{"Elizabeth", "Plath"}
	person3 := Person{"Ronald", "Duthers"}
	person4 := Person{"Lucy", "Brown"}

	people1 := []Person{person1, person2}
	people2 := []Person{person3, person4}

	groups := []Group{
		{"Accounting", people1},
		{"Sales", people2},
	}

	fmt.Println(groups)
	fmt.Println(groups[0])
	fmt.Println(groups[1])
	fmt.Println(groups[1].People[0])
	fmt.Println(groups[1].People[0].LastName)
}
