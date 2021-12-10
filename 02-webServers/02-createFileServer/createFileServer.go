package _030202_createFileServer_

import (
	"fmt"
	"net/http"
)

func Demo() {
	fmt.Println("\n030202 Web Servers: Create File Server")

	fs := http.FileServer(http.Dir("/Users/acdt/lrn/go-web/02-webServers/02-createFileServer"))

	// make files in a dir browsable through the URL
	http.Handle("/projectFiles/", http.StripPrefix("/projectFiles", fs))

	// serve specific file
	http.HandleFunc("/samplepdf/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/Users/acdt/lrn/go-web/02-webServers/02-createFileServer/pdf/sample.pdf")
	})

	//http.ListenAndServe(":9000", nil)

	fmt.Println("Uncomment line 21 in `createWebServer.go` to demo this chapter.")
}
