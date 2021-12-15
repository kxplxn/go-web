package _030206_httprouter_

import "fmt"

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Name: %s %s", p.ByName("firstname"), p.ByName("lastname"))
}

func Demo() {
	fmt.Println("\n030206 Web Servers: Using httprouter")

	//router := httprouter.New()
	//router.GET("/name/:firstname/:lastname", HandleName)
	//
	//http.ListenAndServe(":8080", router)

	fmt.Println("Uncomment line 17-20 in `httprouter.go` to demo this chapter.")
}
