package _030405_variables_

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

func Demo() {
	fmt.Println("\n030405 Go Templates: Variables")

	//tmplFile := "/Users/acdt/lrn/go-web/04-templates/05-variables/personlist.tmpl"
	//
	//http.HandleFunc("/shortlist/", func(w http.ResponseWriter, r *http.Request) {
	//	tmpl, _ := template.ParseFiles(tmplFile)
	//	people := []Person{
	//		{"Bob", "Barker"},
	//		{"Larry", "Flint"},
	//	}
	//	tmpl.Execute(w, people)
	//})
	//
	//http.HandleFunc("/longlist/", func(w http.ResponseWriter, r *http.Request) {
	//	tmpl, _ := template.ParseFiles(tmplFile)
	//	people := []Person{
	//		{"Bob", "Barker"},
	//		{"Larry", "Flint"},
	//		{"Susan", "Sarandon"},
	//		{"Brad", "Pitt"},
	//		{"Betty", "White"},
	//	}
	//	tmpl.Execute(w, people)
	//})
	//
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	tmpl, _ := template.ParseFiles(tmplFile)
	//	tmpl.Execute(w, nil)
	//})
	//
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment lines 15-43 in `flags.go` to demo this chapter.")
}
