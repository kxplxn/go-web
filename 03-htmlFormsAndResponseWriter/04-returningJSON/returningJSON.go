package _03034_returningJSON_

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

type PersonJSONFields struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

type PersonJSONMarshaler struct {
	FirstName string
	LastName  string
}

func (p *PersonJSONMarshaler) MarshalJSON() ([]byte, error) {
	return []byte(`{"Name":"` + p.FirstName + ` ` + p.LastName + `"}`), nil
}

func Demo() {
	fmt.Println("\n030304 HTML Forms & ResponseWriter: Returning JSON")

	//http.HandleFunc("/json/default/", func(w http.ResponseWriter, r *http.Request) {
	//	j, _ := json.Marshal(&Person{FirstName: "Bob", LastName: "Baker"})
	//	w.Write(j)
	//})
	//
	//http.HandleFunc("/json/fields/", func(w http.ResponseWriter, r *http.Request) {
	//	j, _ := json.Marshal(&PersonJSONFields{FirstName: "Bob", LastName: "Baker"})
	//	w.Write(j)
	//})
	//
	//http.HandleFunc("/json/marshaler/", func(w http.ResponseWriter, r *http.Request) {
	//	j, _ := json.Marshal(&PersonJSONMarshaler{FirstName: "Bob", LastName: "Baker"})
	//	w.Write(j)
	//})
	//
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment lines 29-44 in `returningJSON.go` to demo this chapter.")
}
