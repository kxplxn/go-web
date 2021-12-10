package _030201_createWebServer_

import (
	"fmt"
	"net/http"
)

func Demo() {
	fmt.Println("\n030201 Web Servers: Create Web Server")

	// catch-all
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/unrecognized_path.html")
	})

	http.HandleFunc("/demopath1/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath1.html")
	})

	http.HandleFunc("/demopath1/subpatha/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath1_subpatha.html")
	})

	http.HandleFunc("/demopath2", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath2.html")
	})

	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment line 28 in `createWebServer.go` to demo this chapter.")
}
