package _030204_multipleHandlers_

import (
	"fmt"
	"net/http"
)

type PersonHandler struct {
	FirstName string
	LastName  string
}

func (h *PersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Name: %v %v", h.FirstName, h.LastName)
}

type CarHandler struct {
	Make  string
	Model string
	Year  int
}

func (h *CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Make: %v, Model: %v, Year: %v", h.Make, h.Model, h.Year)
}

func Demo() {
	fmt.Println("\n030204 Web Servers: Working with Multiple Handlers")

	http.Handle("/person", &PersonHandler{"Amos", "Kaplan"})
	http.Handle("/car", &CarHandler{"Ford", "Fiesta", 2005})
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment line 32 in `multipleHandlers.go` to demo this chapter.")
}
