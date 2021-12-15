package _030301_htmlForms_

import (
	"fmt"
)

func Demo() {
	fmt.Println("\n030301 HTML Forms & ResponseWriter: HTML Forms")

	//http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
	//	if r.Method == http.MethodPost {
	//		r.ParseForm()
	//		fmt.Println("First Name from Form:", r.Form["firstname"])
	//		fmt.Println("Last Name from Form:", r.Form["lastname"])
	//		fmt.Println("First Name from PostForm:", r.PostForm["firstname"])
	//		fmt.Println("Last Name from PostForm:", r.PostForm["lastname"])
	//	}
	//	http.ServeFile(w, r, "03-htmlFormsAndResponseWriter/01-htmlForms/PersonForm.html")
	//})
	//
	//http.HandleFunc("/file/", func(w http.ResponseWriter, r *http.Request) {
	//	if r.Method == http.MethodPost {
	//		r.ParseMultipartForm(0)
	//		fmt.Println("File information:", r.MultipartForm)
	//		file, _, _ := r.FormFile("file")
	//		bytes, _ := ioutil.ReadAll(file)
	//		fmt.Println("\n\nFile content:", string(bytes))
	//	}
	//	http.ServeFile(w, r, "03-htmlFormsAndResponseWriter/01-htmlForms/FileForm.html")
	//})
	//
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Uncomment line 10-32 in `htmlForms.go` to demo this chapter.")
}
