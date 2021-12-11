package _0302_webServers_

import (
	createWebServer "go-web/02-webServers/01-createWebServer"
	createFileServer "go-web/02-webServers/02-createFileServer"
	handlers "go-web/02-webServers/03-handlers"
	multipleHandlers "go-web/02-webServers/04-multipleHandlers"
	serveMux "go-web/02-webServers/05-serveMux"
	httprouter "go-web/02-webServers/06-httprouter"
)

func Demos() {
	createWebServer.Demo()
	createFileServer.Demo()
	handlers.Demo()
	multipleHandlers.Demo()
	serveMux.Demo()
	httprouter.Demo()
}
