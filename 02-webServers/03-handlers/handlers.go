package _020303_handlers_

import (
	"fmt"
	"net/http"
)

type TestHandler struct{}

func (*TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom handler called TestHandler.")
}

func TestHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom handler function called TestHandlerFunc.")
}

func Demo() {
	fmt.Println("\n030203 Web Servers: Working with Handlers")

	http.Handle("/handler", &TestHandler{})
	http.HandleFunc("/handlerfunc", TestHandlerFunc)
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment line 23 in `handlers.go` to demo this chapter.")
}
