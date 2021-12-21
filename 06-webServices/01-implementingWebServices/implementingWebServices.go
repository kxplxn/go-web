package _030601_implementingWebServices_

import (
	"fmt"
)

func Demo() {
	fmt.Println("\n030601 Web Services: Implementing Web Services in Go")

	//http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusBadRequest)
	//})
	//
	//http.HandleFunc("/withargs/", func (w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(
	//		w,
	//		"This is a response from a %s request. The arguments are: arg1:%s, arg2:%s",
	//		r.Method,
	//		r.URL.Query().Get("arg1"),
	//		r.URL.Query().Get("arg2"),
	//	)
	//})
	//
	//http.HandleFunc("/postonlywithbody/", func (w http.ResponseWriter, r *http.Request) {
	//	if r.Method != http.MethodPost {
	//		w.WriteHeader(http.StatusMethodNotAllowed)
	//		return
	//	}
	//	body, _ := ioutil.ReadAll(r.Body)
	//	fmt.Fprintln(w, string(body))
	//})
	//
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment lines 10-33 in `implementingWebServices.go` to demo this chapter.")
}
