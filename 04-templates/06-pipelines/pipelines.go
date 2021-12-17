package _030406_pipelines_

import (
	"fmt"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
}

func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func Demo() {
	fmt.Println("\n030406 Go Templates: Pipelines")

	//tmplFile := "/Users/acdt/lrn/go-web/04-templates/06-pipelines/personlist.tmpl"
	//
	//fmap := template.FuncMap{
	//	"removeSpaces": removeSpaces,
	//	"capitalize": strings.ToUpper,
	//}
	//
	//http.HandleFunc("/shortlist/", func(w http.ResponseWriter, r *http.Request) {
	//	tmpl, _ := template.New("personlist.tmpl").Funcs(fmap).ParseFiles(tmplFile)
	//	people := []Person{
	//		{"Bo b", "Bark   er"},
	//		{"La rry", "Fli nt"},
	//	}
	//	tmpl.Execute(w, people)
	//})
	//
	//http.HandleFunc("/longlist/", func(w http.ResponseWriter, r *http.Request) {
	//	tmpl, _ := template.New("personlist.tmpl").Funcs(fmap).ParseFiles(tmplFile)
	//	people := []Person{
	//		{"Bob", "Barker"},
	//		{"Lar   ry", "Flint"},
	//		{"Su  san", "Sara  ndon"},
	//		{"B rad", "Pi  tt"},
	//		{"B  etty", "Whi  te"},
	//	}
	//	tmpl.Execute(w, people)
	//})
	//
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	tmpl, _ := template.New("personlist.tmpl").Funcs(fmap).ParseFiles(tmplFile)
	//	tmpl.Execute(w, nil)
	//})
	//
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment lines 20-53 in `pipelines.go` to demo this chapter.")
}
