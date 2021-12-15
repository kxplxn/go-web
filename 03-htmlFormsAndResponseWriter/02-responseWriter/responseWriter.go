package _030302_responseWriter_

import (
	"fmt"
)

func Demo() {
	fmt.Println("\n030302 HTML Forms & ResponseWriter: ResponseWriter")

	//http.HandleFunc("/plaintext/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "text/plain")
	//	io.WriteString(w, "<h1>This is plain text.</h1>")
	//})
	//
	//http.HandleFunc("/html/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "text/html")
	//	io.WriteString(w, "<h1>This is HTML.</h1>")
	//})
	//
	//http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "03-htmlFormsAndResponseWriter/02-responseWriter/static.html")
	//})
	//
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment lines 10-24 in `responseWriter.go` to demo this chapter.")
}
