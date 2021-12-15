package _030205_serveMux_

import (
	"fmt"
	"log"
	"net/http"
)

func customHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom handler.")
}

func logRequests(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("serving request...")
		h.ServeHTTP(w, r)
		log.Println("request is served")
	})
}

func Demo() {
	fmt.Println("\n030205 Web Servers: Using ServeMux and DefaultServeMux")

	//serveMux := http.NewServeMux()
	//serveMux.HandleFunc("/custom/", customHandler)

	// wrapping a ServeMux (chaining handlers)
	//loggedServeMux := logRequests(serveMux)

	//http.ListenAndServe(":8080", loggedServeMux)

	fmt.Println("Uncomment line 24-30 in `serveMux.go` to demo this chapter.")
}
