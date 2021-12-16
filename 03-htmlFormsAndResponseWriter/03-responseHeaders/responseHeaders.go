package _030303_responseHeaders_

import (
	"fmt"
)

func Demo() {
	fmt.Println("\n030303 HTML Forms & ResponseWriter: Response Headers")

	//http.HandleFunc("/notfound/", func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(404)
	//})
	//
	//http.HandleFunc("/servererror/", func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(500)
	//})
	//
	//http.HandleFunc("/customheader/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Custom-Header", "This is my custom header.")
	//	io.WriteString(w, "Check out the custom header by pressing F12.")
	//})
	//
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment lines 10-23 in `responseHeaders.go` to demo this chapter.")
}
