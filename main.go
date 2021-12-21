package main

import (
	basics "go-web/01-basics"
	webServers "go-web/02-webServers"
	htmlFormsAndResponseWriter "go-web/03-htmlFormsAndResponseWriter"
	templates "go-web/04-templates"
	dataStorage "go-web/05-dataStorage"
	webServices "go-web/06-webServices"
)

func main() {
	basics.Demos()
	webServers.Demos()
	htmlFormsAndResponseWriter.Demos()
	templates.Demos()
	dataStorage.Demos()
	webServices.Demos()
}
