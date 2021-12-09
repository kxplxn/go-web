package main

import (
	demo "go-web/01-basics/03-structuringApps"
	"go-web/01-basics/03-structuringApps/html"
	"go-web/01-basics/03-structuringApps/plaintext"
	"net/http"
)

func main() {
	http.Handle("/person/html", &html.PersonHandler{Person: demo.Person{FirstName: "Bob", LastName: "Smith"}})
	http.Handle("/person/plaintext", &plaintext.PersonHandler{Person: demo.Person{FirstName: "Bob", LastName: "Smith"}})

	http.Handle("/car/html", &html.CarHandler{Car: demo.Car{Make: "Ford", Model: "Fiesta", Year: 2005}})
	http.Handle("/car/plaintext", &plaintext.CarHandler{Car: demo.Car{Make: "Ford", Model: "Fiesta", Year: 2005}})

	http.ListenAndServe(":8080", nil)
}
