package _030603_workingWithJSON_

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string   `json:"fname"`
	LastName  string   `json:"lname"`
	Hobbies   []string `json:"hobbies"`
}

func Demo() {
	fmt.Println("\n030603 Web Services: Working with JSON")

	joe := Person{FirstName: "Joe", LastName: "Smith"}
	joe.Hobbies = []string{"Skiing", "Wind Surfing"}

	jsonOutput, _ := json.MarshalIndent(&joe, " ", "   ")
	fmt.Println(string(jsonOutput))

	joeFromJSON := Person{}
	json.Unmarshal(jsonOutput, &joeFromJSON)
	fmt.Println(joeFromJSON)
}
