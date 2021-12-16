package _030402_actions_

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

func Demo() {
	fmt.Println("\n030402 Go Templates: Actions")

	//templateFile := "/Users/acdt/lrn/go-web/04-templates/02-actions/personlist.tmpl"
	//
	//http.HandleFunc("/shortlist/", func(w http.ResponseWriter, r *http.Request) {
	//	t, _ := template.ParseFiles(templateFile)
	//	p := []Person{
	//		{"Bob", "Barker"},
	//		{"Larry", "Flint"},
	//	}
	//	t.Execute(w, p)
	//})
	//
	//http.HandleFunc("/longlist/", func(w http.ResponseWriter, r *http.Request) {
	//	t, _ := template.ParseFiles(templateFile)
	//	p := []Person{
	//		{"Bob", "Barker"},
	//		{"Larry", "Flint"},
	//		{"Susan", "Sarandon"},
	//		{"Brad", "Pitt"},
	//		{"Betty", "White"},
	//	}
	//	t.Execute(w, p)
	//})
	//
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	t, _ := template.ParseFiles(templateFile)
	//	t.Execute(w, nil)
	//})
	//
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment lines 15-43 in `actions.go` to demo this chapter.")
}
