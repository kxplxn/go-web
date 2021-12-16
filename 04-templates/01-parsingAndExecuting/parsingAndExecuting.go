package _030401_parsingAndExecuting_

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

func Demo() {
	fmt.Println("\n030401 Go Templates: Parsing and Executing")

	//tmpl := `<b><font color="green">First Name: </font></b> {{ .FirstName }} <br/>
	//	<b><font color="red">Last Name: </font></b> {{ .LastName }}`
	//
	//router := httprouter.New()
	//
	//router.GET("/name/:firstName/:lastName/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//	t := template.New("NameTemplate")
	//	tp, _ := t.Parse(tmpl)
	//	tp.Execute(w, &Person{ps.ByName("firstName"), ps.ByName("lastName")})
	//})
	//
	//http.ListenAndServe(":8080", router)

	fmt.Println("Uncomment lines 15-26 in `parsingAndExecuting.go` to demo this chapter.")
}
