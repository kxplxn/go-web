package _030407_functions_

import (
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
}

func Demo() {
	fmt.Println("\n030407 Go Templates: Functions")

	//tmplFile := "/Users/acdt/lrn/go-web/04-templates/07-functions/personlist.tmpl"
	//
	//http.HandleFunc("/longlist/", func(w http.ResponseWriter, r *http.Request) {
	//	tmpl, _ := template.ParseFiles(tmplFile)
	//	people := []person{
	//		{"Bob", "Barker"},
	//		{"Larry", "Flint"},
	//		{"Susan", "Sarandon"},
	//		{"Brad", "Pitt"},
	//		{"Betty", "White"},
	//	}
	//	tmpl.Execute(w, people)
	//})
	//
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment lines 15-29 in `functions.go` to demo this chapter.")
}
